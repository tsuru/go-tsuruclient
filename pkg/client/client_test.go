package client

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tsuru/go-tsuruclient/pkg/tsuru"
)

func overrideHome(t *testing.T) (string, func()) {
	os.Unsetenv("TSURU_TARGET")
	os.Unsetenv("TSURU_TOKEN")
	tmpDir, err := ioutil.TempDir("", "tsuru")
	require.Nil(t, err)
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	return tmpDir, func() {
		os.Setenv("HOME", oldHome)
		os.RemoveAll(tmpDir)
	}
}

func Test_ClientFromEnvironment(t *testing.T) {
	tmpHome, rollback := overrideHome(t)
	defer rollback()
	var lastReq *http.Request
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lastReq = r
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}))
	defer srv.Close()
	t.Run("server from env", func(t *testing.T) {
		defer os.Unsetenv("TSURU_TARGET")
		defer os.Unsetenv("TSURU_TOKEN")
		os.Setenv("TSURU_TARGET", srv.URL)
		os.Setenv("TSURU_TOKEN", "mytoken")
		cli, err := ClientFromEnvironment(nil)
		require.Nil(t, err)
		obj, rsp, err := cli.AppApi.AppList(nil, nil)
		require.Nil(t, err)
		assert.Equal(t, http.StatusOK, rsp.StatusCode)
		assert.Equal(t, []tsuru.MiniApp{}, obj)
		assert.Equal(t, "bearer mytoken", lastReq.Header.Get("Authorization"))
	})
	t.Run("server from files", func(t *testing.T) {
		dir := filepath.Join(tmpHome, ".tsuru")
		err := ioutil.WriteFile(filepath.Join(dir, "target"), []byte(srv.URL), 0600)
		require.Nil(t, err)
		err = ioutil.WriteFile(filepath.Join(dir, "token"), []byte("mytokenfile"), 0600)
		require.Nil(t, err)
		cli, err := ClientFromEnvironment(nil)
		require.Nil(t, err)
		obj, rsp, err := cli.AppApi.AppList(nil, nil)
		require.Nil(t, err)
		assert.Equal(t, http.StatusOK, rsp.StatusCode)
		assert.Equal(t, []tsuru.MiniApp{}, obj)
		assert.Equal(t, "bearer mytokenfile", lastReq.Header.Get("Authorization"))
	})
}
