/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// Dynamic router
type DynamicRouter struct {
	Name           string                 `json:"name,omitempty"`
	Type           string                 `json:"type,omitempty"`
	ReadinessGates []string               `json:"readinessGates,omitempty"`
	Config         map[string]interface{} `json:"config,omitempty"`
}
