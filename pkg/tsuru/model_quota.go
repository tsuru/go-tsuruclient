/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type Quota struct {
	Inuse int64 `json:"inuse,omitempty"`
	Limit int64 `json:"limit,omitempty"`
}
