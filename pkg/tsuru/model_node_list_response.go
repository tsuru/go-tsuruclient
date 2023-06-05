/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.15
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type NodeListResponse struct {
	Nodes    []Node    `json:"nodes,omitempty"`
	Machines []Machine `json:"machines,omitempty"`
}
