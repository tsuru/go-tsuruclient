/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type TeamInfo struct {
	Name  string   `json:"name,omitempty"`
	Tags  []string `json:"tags,omitempty"`
	Users []User   `json:"users,omitempty"`
	Pools []Pool   `json:"pools,omitempty"`
	Apps  []App    `json:"apps,omitempty"`
}
