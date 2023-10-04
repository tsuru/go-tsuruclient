/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.18
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type ServiceInstanceBind struct {
	NoRestart  bool              `json:"noRestart,omitempty"`
	Parameters map[string]string `json:"parameters,omitempty"`
}
