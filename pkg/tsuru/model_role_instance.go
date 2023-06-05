/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.15
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// Association between a role and a context value.
type RoleInstance struct {
	Name         string `json:"name,omitempty"`
	Contextvalue string `json:"contextvalue,omitempty"`
}
