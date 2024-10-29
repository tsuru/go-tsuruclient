/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// App plan override.
type PlanOverride struct {
	Memory   *int64   `json:"memory,omitempty"`
	Cpumilli *int     `json:"cpumilli,omitempty"`
	CpuBurst *float64 `json:"cpuBurst,omitempty"`
}
