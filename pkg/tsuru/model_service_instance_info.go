/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.19
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type ServiceInstanceInfo struct {
	Apps            []string          `json:"apps,omitempty"`
	Jobs            []string          `json:"jobs,omitempty"`
	Teams           []string          `json:"teams,omitempty"`
	Teamowner       string            `json:"teamowner,omitempty"`
	Description     string            `json:"description,omitempty"`
	Pool            string            `json:"pool,omitempty"`
	Planname        string            `json:"planname,omitempty"`
	Plandescription string            `json:"plandescription,omitempty"`
	Tags            []string          `json:"tags,omitempty"`
	Custominfo      map[string]string `json:"custominfo,omitempty"`
	Parameters      map[string]string `json:"parameters,omitempty"`
}
