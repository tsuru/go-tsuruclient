/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.22
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// Assign a role
type RoleAssignData struct {
	Name         string `json:"name,omitempty"`
	Contextvalue string `json:"contextvalue,omitempty"`
	Roletarget   string `json:"roletarget,omitempty"`
	Sufix        string `json:"sufix,omitempty"`
	Version      string `json:"version,omitempty"`
}
