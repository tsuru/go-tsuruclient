// Copyright 2023 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/tsuru/go-tsuruclient/pkg/client"
)

func TestTsuruProxyTransport(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		transport   func(*TsuruProxyTransport) *TsuruProxyTransport
		request     func() *http.Request
		assert      func(t *testing.T, r *http.Request)
		expectedErr string
	}{
		"missing Tsuru target": {
			transport: func(_ *TsuruProxyTransport) *TsuruProxyTransport {
				return &TsuruProxyTransport{}
			},
			request: func() *http.Request {
				r, err := http.NewRequest("GET", "http://example.com/", nil)
				require.NoError(t, err)
				return r
			},
			expectedErr: "Get \"http://example.com/\": Tsuru target not provided",
		},

		"missing Tsuru token": {
			transport: func(_ *TsuruProxyTransport) *TsuruProxyTransport {
				return &TsuruProxyTransport{
					Target: "http://example.com",
				}
			},
			request: func() *http.Request {
				r, err := http.NewRequest("GET", "http://example.com/", nil)
				require.NoError(t, err)
				return r
			},
			expectedErr: "Get \"http://example.com/\": Tsuru token not provided",
		},

		"missing Tsuru Service": {
			transport: func(_ *TsuruProxyTransport) *TsuruProxyTransport {
				return &TsuruProxyTransport{
					Target: "http://example.com",
					Token:  "my-fake-token",
				}
			},
			request: func() *http.Request {
				r, err := http.NewRequest("GET", "http://example.com/", nil)
				require.NoError(t, err)
				return r
			},
			expectedErr: "Get \"http://example.com/\": Tsuru service name not provided",
		},

		"missing Tsuru instance": {
			transport: func(_ *TsuruProxyTransport) *TsuruProxyTransport {
				return &TsuruProxyTransport{
					Target:  "http://example.com",
					Token:   "my-fake-token",
					Service: "my-service",
				}
			},
			request: func() *http.Request {
				r, err := http.NewRequest("GET", "http://example.com/", nil)
				require.NoError(t, err)
				return r
			},
			expectedErr: "Get \"http://example.com/\": Tsuru instance name not provided",
		},

		"with additional query params": {
			request: func() *http.Request {
				r, err := http.NewRequest("GET", "http://example.com/resources/my-instance/info?q=1&w=2&e=3", nil)
				require.NoError(t, err)
				return r
			},
		},

		"with custom headers": {
			request: func() *http.Request {
				r, err := http.NewRequest("GET", "http://example.com/resources/my-instance/info", nil)
				require.NoError(t, err)

				r.Header.Set("X-FooBar", "bazzzz")
				r.Header.Add("X-Header", "v1")
				r.Header.Add("X-Header", "v2")

				return r
			},
			assert: func(t *testing.T, r *http.Request) {
				assert.Equal(t, "bazzzz", r.Header.Get("X-FooBar"))
				assert.Equal(t, []string{"v1", "v2"}, r.Header["X-Header"])
			},
		},

		"with body content": {
			request: func() *http.Request {
				r, err := http.NewRequest("POST", "http://example.com/resources/my-instance/autoscale", strings.NewReader("hello world!"))
				require.NoError(t, err)
				return r
			},
			assert: func(t *testing.T, r *http.Request) {
				body, err := io.ReadAll(r.Body)
				require.NoError(t, err)
				assert.Equal(t, "hello world!", string(body))
			},
		},

		"with custom base transport": {
			transport: func(tpt *TsuruProxyTransport) *TsuruProxyTransport {
				tpt.Base = &dummyTransport{}
				return tpt
			},
			request: func() *http.Request {
				r, err := http.NewRequest("GET", "http://example.com/resources/my-instance/autoscale", nil)
				require.NoError(t, err)
				return r
			},
			assert: func(t *testing.T, r *http.Request) {
				assert.Equal(t, "1", r.Header.Get("X-Dummy-Transport"))
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			require.NotNil(t, tt.request, "you must provide a request generator function")

			wanted := tt.request()

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, got *http.Request) {
				assert.Equal(t, wanted.Method, got.Method)
				assert.Equal(t, "/1.0/services/my-service/proxy/my-instance", got.URL.Path)
				assert.Equal(t, "Bearer f4k3-tsuru-t0k3n", got.Header.Get("Authorization"))

				if wanted.URL.RawQuery == "" {
					assert.Equal(t, fmt.Sprintf("callback=%s", wanted.URL.Path), got.URL.RawQuery)
				} else {
					assert.Equal(t, fmt.Sprintf("callback=%s&%s", wanted.URL.Path, wanted.URL.RawQuery), got.URL.RawQuery)
				}

				if tt.assert != nil {
					tt.assert(t, got)
				}
			}))
			defer server.Close()

			transport := &TsuruProxyTransport{
				Target:   server.URL,
				Token:    "f4k3-tsuru-t0k3n",
				Service:  "my-service",
				Instance: "my-instance",
			}

			if tt.transport != nil {
				transport = tt.transport(transport)
			}

			c := &http.Client{Transport: transport}
			_, err := c.Do(wanted)
			if tt.expectedErr == "" {
				require.NoError(t, err)
				return
			}

			assert.EqualError(t, err, tt.expectedErr)
		})
	}
}

type dummyTransport struct{}

func (d *dummyTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("X-Dummy-Transport", "1")
	return http.DefaultTransport.RoundTrip(r)
}
