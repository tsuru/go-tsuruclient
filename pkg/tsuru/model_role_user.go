/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// Role of an user.
type RoleUser struct {
	Name         string `json:"name,omitempty"`
	Contexttype  string `json:"contexttype,omitempty"`
	Contextvalue string `json:"contextvalue,omitempty"`
	Group        string `json:"group,omitempty"`
}
