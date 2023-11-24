/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.19
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// Add a permission
type PermissionData struct {
	Name       string   `json:"name,omitempty"`
	Permission []string `json:"permission,omitempty"`
}
