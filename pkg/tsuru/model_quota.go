/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// Quota struct for Quota
type Quota struct {
	Inuse int64 `json:"inuse,omitempty"`
	Limit int64 `json:"limit,omitempty"`
}
