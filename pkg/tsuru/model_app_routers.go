/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.22
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type AppRouters struct {
	Addresses []string               `json:"addresses,omitempty"`
	Opts      map[string]interface{} `json:"opts,omitempty"`
	Name      string                 `json:"name,omitempty"`
}
