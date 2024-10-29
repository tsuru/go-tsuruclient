/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type Pool struct {
	Name        string              `json:"name,omitempty"`
	Default     bool                `json:"default,omitempty"`
	Provisioner string              `json:"provisioner,omitempty"`
	Public      bool                `json:"public,omitempty"`
	Teams       []string            `json:"teams,omitempty"`
	Allowed     map[string][]string `json:"allowed,omitempty"`
	Labels      map[string]string   `json:"labels,omitempty"`
}
