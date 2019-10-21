/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru
// NodeGetResponse struct for NodeGetResponse
type NodeGetResponse struct {
	Node Node `json:"node,omitempty"`
	Status NodeStatus `json:"status,omitempty"`
	Units Unit `json:"units,omitempty"`
}
