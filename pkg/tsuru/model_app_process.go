/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.19
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type AppProcess struct {
	Name     string   `json:"name,omitempty"`
	Plan     string   `json:"plan,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`
}
