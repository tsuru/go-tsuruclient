/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// App Run options
type AppRunOpts struct {
	Once     bool   `json:"once,omitempty"`
	Isolated bool   `json:"isolated,omitempty"`
	Command  string `json:"command,omitempty"`
}
