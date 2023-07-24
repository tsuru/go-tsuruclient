package tsurutests

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tsuru/go-tsuruclient/pkg/tsuru"
)

// This package contains tests to validate changes to templates made for this
// project. If these tests pass without custom templates it's assumed that we
// don't need our custom templates anymore.

func Test_MultiQueryParams(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.RawQuery
		assert.Equal(t, "env=e1&env=e2&norestart=false", query)
	}))
	defer srv.Close()

	cli := tsuru.NewAPIClient(&tsuru.Configuration{
		BasePath: srv.URL,
	})
	_, err := cli.AppApi.EnvUnset(context.Background(), "myapp", []string{"e1", "e2"}, false)
	require.NoError(t, err)
}

func Test_ByteTypes(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		require.NoError(t, err)
		assert.Equal(t, `{"cacert":"YWJj"}`+"\n", string(data))
	}))
	defer srv.Close()

	cli := tsuru.NewAPIClient(&tsuru.Configuration{
		BasePath: srv.URL,
	})
	_, err := cli.ClusterApi.ClusterCreate(context.Background(), tsuru.Cluster{
		Cacert: []byte("abc"),
	})
	require.NoError(t, err)
}

func Test_DoNotReadBodyForStreamingRequests(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("lazy response"))
	}))
	defer srv.Close()

	cli := tsuru.NewAPIClient(&tsuru.Configuration{
		BasePath: srv.URL,
	})
	rsp, err := cli.AppApi.AppRestart(context.Background(), "myapp", tsuru.AppStartStop{})
	require.NoError(t, err)
	data, err := ioutil.ReadAll(rsp.Body)
	require.NoError(t, err)
	assert.Equal(t, "lazy response", string(data))
}

func Test_AcceptPlainTextErrorResponses(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("my error"))
	}))
	defer srv.Close()

	cli := tsuru.NewAPIClient(&tsuru.Configuration{
		BasePath: srv.URL,
	})
	_, err := cli.AppApi.AppRestart(context.Background(), "myapp", tsuru.AppStartStop{})
	require.Error(t, err)
	apiErr, ok := err.(tsuru.GenericOpenAPIError)
	require.Equal(t, true, ok)
	assert.Equal(t, "my error", string(apiErr.Body()))
}

func Test_HasStatusCodeInResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	cli := tsuru.NewAPIClient(&tsuru.Configuration{
		BasePath: srv.URL,
	})
	_, err := cli.AppApi.AppRestart(context.Background(), "myapp", tsuru.AppStartStop{})
	require.Error(t, err)
	apiErr, ok := err.(tsuru.GenericOpenAPIError)
	require.Equal(t, true, ok)
	assert.Equal(t, http.StatusNotFound, apiErr.StatusCode())
}

func Test_AcceptPlainTextRegularResponses(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("my result"))
	}))
	defer srv.Close()

	cli := tsuru.NewAPIClient(&tsuru.Configuration{
		BasePath: srv.URL,
	})
	rsp, _, err := cli.UserApi.APITokenGet(context.Background(), "me@me.com")
	require.NoError(t, err)
	assert.Equal(t, "my result", rsp)
}
