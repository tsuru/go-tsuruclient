/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 */

// NOTE: Unlike most files in this package, the current one is not self-generated.
// PLEASE DO NOT REMOVE IT.

package tsuru

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServiceApiService_InstanceProxy(t *testing.T) {
	var called bool

	tests := []struct {
		service  string
		instance string
		pr       ServiceInstanceProxyRequest
		handler  http.Handler
		assert   func(t *testing.T, r *http.Response, e error)
	}{
		{
			assert: func(t *testing.T, r *http.Response, e error) {
				assert.Nil(t, r)
				assert.Error(t, e)
				assert.EqualError(t, e, "service cannot be empty")
			},
		},
		{
			service: "service-name",
			assert: func(t *testing.T, r *http.Response, e error) {
				assert.Nil(t, r)
				assert.Error(t, e)
				assert.EqualError(t, e, "instance cannot be empty")
			},
		},
		{
			service:  "service-name",
			instance: "instance-name",
			assert: func(t *testing.T, r *http.Response, e error) {
				assert.Nil(t, r)
				assert.Error(t, e)
				assert.EqualError(t, e, "callback path cannot be empty")
			},
		},
		{
			service:  "service-name",
			instance: "instance-name",
			pr: ServiceInstanceProxyRequest{
				Callback: "/path/to/resource",
			},
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				called = true
				assert.Equal(t, "GET", r.Method)
				assert.Equal(t, "/1.0/services/service-name/proxy/instance-name", r.URL.Path)
				assert.Equal(t, "callback=%2Fpath%2Fto%2Fresource", r.URL.RawQuery)
				fmt.Fprintf(w, "some instance response")
			}),
			assert: func(t *testing.T, r *http.Response, e error) {
				require.NoError(t, e)
				require.True(t, called)
				require.NotNil(t, r)
				rawBody, err := ioutil.ReadAll(r.Body)
				require.NoError(t, err)
				defer r.Body.Close()
				assert.Equal(t, "some instance response", string(rawBody))
			},
		},
		{
			service:  "service-name",
			instance: "instance-name",
			pr: ServiceInstanceProxyRequest{
				Method:   http.MethodPost,
				Callback: "/path/to/resource",
				Header: map[string]string{
					"Content-Type":    "application/x-www-form-urlencoded",
					"X-Custom-Header": "Value",
				},
				Body: ioutil.NopCloser(strings.NewReader("foo=bar")),
			},
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				called = true
				assert.Equal(t, "POST", r.Method)
				assert.Equal(t, "/1.0/services/service-name/proxy/instance-name", r.URL.Path)
				assert.Equal(t, "callback=%2Fpath%2Fto%2Fresource", r.URL.RawQuery)
				assert.Equal(t, "application/x-www-form-urlencoded", r.Header.Get("Content-Type"))
				assert.Equal(t, "Value", r.Header.Get("X-Custom-Header"))

				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "some error response")
			}),
			assert: func(t *testing.T, r *http.Response, e error) {
				require.NoError(t, e)
				require.True(t, called)
				require.NotNil(t, r)
				rawBody, err := ioutil.ReadAll(r.Body)
				require.NoError(t, err)
				defer r.Body.Close()
				assert.Equal(t, http.StatusInternalServerError, r.StatusCode)
				assert.Equal(t, "some error response", string(rawBody))
			},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			called = false

			srv := httptest.NewServer(tt.handler)
			defer srv.Close()

			a := &ServiceApiService{
				client: &APIClient{
					cfg: &Configuration{BasePath: srv.URL, HTTPClient: http.DefaultClient},
				},
			}

			r, err := a.InstanceProxy(context.TODO(), tt.service, tt.instance, tt.pr)
			require.NotNil(t, tt.assert)
			tt.assert(t, r, err)
		})
	}
}
