/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type TeamGroup struct {
	Group string   `json:"group,omitempty"`
	Roles []string `json:"roles,omitempty"`
}
