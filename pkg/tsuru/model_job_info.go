/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.18
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type JobInfo struct {
	Job   Job    `json:"job,omitempty"`
	Units []Unit `json:"units,omitempty"`
	// Service instance binds on the job
	ServiceInstanceBinds []AppServiceInstanceBinds `json:"serviceInstanceBinds,omitempty"`
}
