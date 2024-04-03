package client

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/tsuru/go-tsuruclient/pkg/config"
	"golang.org/x/oauth2"
)

func NewOIDCTokenSource(tokenV2 *config.TokenV2) oauth2.TokenSource {
	baseTokenSource := tokenV2.OAuth2Config.TokenSource(context.Background(), tokenV2.OAuth2Token)
	return newTokenSourceFSStorage(baseTokenSource, tokenV2)
}

type TokenSourceFSStorage struct {
	BaseTokenSource oauth2.TokenSource
	LastToken       *config.TokenV2
}

var _ oauth2.TokenSource = &TokenSourceFSStorage{}

func (t *TokenSourceFSStorage) Token() (*oauth2.Token, error) {
	newToken, err := t.BaseTokenSource.Token()
	if err != nil {
		return nil, err
	}

	if !reflect.DeepEqual(t.LastToken.OAuth2Token, newToken) {
		fmt.Fprintf(os.Stderr, "The OIDC token was refreshed and expiry in %s\n", time.Since(newToken.Expiry)*-1)

		t.LastToken.OAuth2Token = newToken
		err = config.WriteTokenV2(*t.LastToken)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not write refreshed token: %s\n", err.Error())
			return nil, err
		}

		err = config.WriteTokenV1(newToken.AccessToken)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not write legacy refreshed token: %s\n", err.Error())
			return nil, err
		}
	}

	return newToken, nil
}

func newTokenSourceFSStorage(baseTokenSource oauth2.TokenSource, tokenV2 *config.TokenV2) oauth2.TokenSource {
	return &TokenSourceFSStorage{
		BaseTokenSource: baseTokenSource,
		LastToken:       tokenV2,
	}
}

type OIDCTokenProvider struct {
	OAuthTokenSource oauth2.TokenSource
}

func (ts *OIDCTokenProvider) Token() (string, error) {
	t, err := ts.OAuthTokenSource.Token()
	if err != nil {
		return "", err
	}
	return t.AccessToken, nil
}
