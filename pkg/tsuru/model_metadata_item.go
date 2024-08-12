/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.22
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// Metadata items
type MetadataItem struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Delete bool   `json:"delete,omitempty"`
}
