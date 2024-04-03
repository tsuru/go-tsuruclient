package client

import (
	"net/http"

	"github.com/tsuru/go-tsuruclient/pkg/config"
	"github.com/tsuru/go-tsuruclient/pkg/tsuru"
	"golang.org/x/oauth2"
)

func ClientFromEnvironment(cfg *tsuru.Configuration) (*tsuru.APIClient, error) {
	if cfg == nil {
		cfg = &tsuru.Configuration{}
	}
	var err error
	if cfg.BasePath == "" {
		cfg.BasePath, err = config.GetTarget()
		if err != nil {
			return nil, err
		}
	}
	if cfg.DefaultHeader == nil {
		cfg.DefaultHeader = map[string]string{}
	}
	if _, authSet := cfg.DefaultHeader["Authorization"]; !authSet {
		tokenV2, err := config.ReadTokenV2()
		if err != nil {
			return nil, err
		}

		teamToken := config.ReadTeamToken()
		if tokenV2 != nil && tokenV2.Scheme == "oidc" && teamToken == "" {
			oidcTokenSource := NewOIDCTokenSource(tokenV2)
			config.DefaultTokenProvider = &OIDCTokenProvider{OAuthTokenSource: oidcTokenSource}

			cfg.HTTPClient = &http.Client{
				Transport: &oauth2.Transport{
					Base:   http.DefaultTransport,
					Source: oidcTokenSource,
				},
			}
		} else if token, tokenErr := config.ReadTokenV1(); tokenErr == nil && token != "" {
			cfg.DefaultHeader["Authorization"] = "bearer " + token
		}
	}
	cli := tsuru.NewAPIClient(cfg)
	return cli, nil
}

func RoundTripperAndTokenProvider() (http.RoundTripper, config.TokenProvider, error) {
	tokenV2, err := config.ReadTokenV2()
	if err != nil {
		return nil, nil, err
	}

	teamToken := config.ReadTeamToken()
	if tokenV2 != nil && tokenV2.Scheme == "oidc" && teamToken == "" {
		oidcTokenSource := NewOIDCTokenSource(tokenV2)
		tokenProvider := &OIDCTokenProvider{OAuthTokenSource: oidcTokenSource}

		roundTripper := &oauth2.Transport{
			Base:   http.DefaultTransport,
			Source: oidcTokenSource,
		}

		return roundTripper, tokenProvider, nil
	}

	return NewTokenV1RoundTripper(), config.TokenProviderV1(), nil
}
