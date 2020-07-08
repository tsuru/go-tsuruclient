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
	"io"
	"net/http"
)

// ServiceInstanceProxyRequest defines the HTTP request to be proxied to the
// Service API.
type ServiceInstanceProxyRequest struct {
	Method   string
	Callback string
	Body     io.ReadCloser
	Header   map[string]string
}

// InstanceProxy calls a Tsuru API endpoint that performs a reverse proxy
// between client and Service API.
func (a *ServiceApiService) InstanceProxy(ctx context.Context, service, instance string, r ServiceInstanceProxyRequest) (*http.Response, error) {
	if service == "" {
		return nil, reportError("service cannot be empty")
	}

	if instance == "" {
		return nil, reportError("instance cannot be empty")
	}

	if r.Callback == "" {
		return nil, reportError("callback path cannot be empty")
	}

	method := http.MethodGet
	if r.Method != "" {
		method = r.Method
	}

	path := fmt.Sprintf("%s/1.0/services/%s/proxy/%s?callback=%s", a.client.cfg.BasePath, service, instance, r.Callback)

	request, err := a.client.prepareRequest(ctx, path, method, nil, r.Header, nil, nil, "", "", nil)
	if err != nil {
		return nil, err
	}
	request.Body = r.Body

	return a.client.callAPI(request)
}
