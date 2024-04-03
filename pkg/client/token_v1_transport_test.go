package client

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tsuru/go-tsuruclient/pkg/config"
	"github.com/tsuru/tsuru/cmd/cmdtest"
	"github.com/tsuru/tsuru/fs/fstest"
)

func TestTokenV1RoundTripperShouldIncludeTheHeaderAuthorizationWhenTsuruTokenFileExists(t *testing.T) {
	os.Unsetenv("TSURU_TOKEN")
	config.SetFileSystem(&fstest.RecordingFs{FileContent: "mytoken"})
	targetInit()
	defer func() {
		config.ResetFileSystem()
	}()
	request, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)
	trans := cmdtest.Transport{Message: "", Status: http.StatusOK}
	rt := &TokenV1RoundTripper{RoundTripper: &trans}
	_, err = rt.RoundTrip(request)
	require.NoError(t, err)
	assert.Equal(t, "bearer mytoken", request.Header.Get("Authorization"))
}

func targetInit() {
	f, _ := config.Filesystem().Create(config.JoinWithUserDir(".tsuru", "target"))
	f.Write([]byte("http://localhost"))
	f.Close()
	config.WriteOnTargetList("test", "http://localhost")
}
