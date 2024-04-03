package client

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tsuru/go-tsuruclient/pkg/config"
	tsuruerr "github.com/tsuru/tsuru/errors"
)

var errUnauthorized = &tsuruerr.HTTP{Code: http.StatusUnauthorized, Message: "unauthorized"}

type TokenV1RoundTripper struct {
	http.RoundTripper
}

func (v *TokenV1RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	roundTripper := v.RoundTripper
	if roundTripper == nil {
		roundTripper = http.DefaultTransport
	}

	if token, err := config.ReadTokenV1(); err == nil && token != "" {
		req.Header.Set("Authorization", "bearer "+token)
	}

	response, err := roundTripper.RoundTrip(req)

	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusUnauthorized {
		if teamToken := config.ReadTeamToken(); teamToken != "" {
			fmt.Fprintln(os.Stderr, "Invalid session - maybe invalid defined token on TSURU_TOKEN envvar")
		} else {
			fmt.Fprintln(os.Stderr, "Invalid session")
		}

		return nil, errUnauthorized
	}

	return response, nil
}

func NewTokenV1RoundTripper() http.RoundTripper {
	return &TokenV1RoundTripper{
		RoundTripper: http.DefaultTransport,
	}
}
