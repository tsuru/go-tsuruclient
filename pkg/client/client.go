package client

import (
	"github.com/tsuru/go-tsuruclient/pkg/config"
	"github.com/tsuru/go-tsuruclient/pkg/tsuru"
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
		// TODO: use transport instead of token v1
		if token, tokenErr := config.ReadTokenV1(); tokenErr == nil && token != "" {
			cfg.DefaultHeader["Authorization"] = "bearer " + token
		}
	}
	cli := tsuru.NewAPIClient(cfg)
	return cli, nil
}
