/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.18
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type Team struct {
	Name        string   `json:"name,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}
