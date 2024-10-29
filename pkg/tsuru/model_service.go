/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type Service struct {
	Id           string `json:"id,omitempty"`
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	Endpoint     string `json:"endpoint,omitempty"`
	MultiCluster string `json:"multi-cluster,omitempty"`
	Team         string `json:"team,omitempty"`
}
