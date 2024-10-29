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
type RoleUpdateData struct {
	Name        string `json:"name,omitempty"`
	ContextType string `json:"contextType,omitempty"`
	NewName     string `json:"newName,omitempty"`
	Description string `json:"description,omitempty"`
}
