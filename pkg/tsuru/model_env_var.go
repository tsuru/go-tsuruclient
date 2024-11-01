/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// Environment variable.
type EnvVar struct {
	Name      string `json:"name,omitempty"`
	Value     string `json:"value,omitempty"`
	Alias     string `json:"alias,omitempty"`
	Public    bool   `json:"public,omitempty"`
	ManagedBy string `json:"managedBy,omitempty"`
}
