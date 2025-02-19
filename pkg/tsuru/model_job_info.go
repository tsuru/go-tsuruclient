/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type JobInfo struct {
	Cluster      string `json:"cluster,omitempty"`
	Job          Job    `json:"job,omitempty"`
	Units        []Unit `json:"units,omitempty"`
	DashboardURL string `json:"dashboardURL,omitempty"`
	// Service instance binds on the job
	ServiceInstanceBinds []AppServiceInstanceBinds `json:"serviceInstanceBinds,omitempty"`
}
