package client

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tsuru/go-tsuruclient/pkg/config"
	"github.com/tsuru/tsuru/fs/fstest"
	"golang.org/x/oauth2"
)

type fakeTokenSource struct{}

func (f *fakeTokenSource) Token() (*oauth2.Token, error) {
	return &oauth2.Token{
		AccessToken: "access-token-321",
		Expiry:      time.Now().Add(time.Hour),
	}, nil
}
func TestTokenSourceFSStorage(t *testing.T) {

	config.SetFileSystem(&fstest.RecordingFs{})

	defer func() {
		config.ResetFileSystem()
	}()

	fts := &fakeTokenSource{}
	tokenSourceFSStorage := &TokenSourceFSStorage{
		BaseTokenSource: fts,
		LastToken: &config.TokenV2{
			OAuth2Token: &oauth2.Token{
				AccessToken: "access-token-123",
			},
		},
	}

	token, err := tokenSourceFSStorage.Token()
	require.NoError(t, err)

	assert.Equal(t, "access-token-321", token.AccessToken)

	tokenV1fromConfig, err := config.ReadTokenV1()
	require.NoError(t, err)
	assert.Equal(t, "access-token-321", tokenV1fromConfig)

	tokenV2fromConfig, err := config.ReadTokenV2()
	require.NoError(t, err)
	assert.Equal(t, "access-token-321", tokenV2fromConfig.OAuth2Token.AccessToken)

}
