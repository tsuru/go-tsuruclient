/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type NodeCheckResult struct {
	Name       string `json:"name,omitempty"`
	Err        string `json:"err,omitempty"`
	Successful bool   `json:"successful,omitempty"`
}
