/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.18
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type JobSpec struct {
	Container JobSpecContainer `json:"container,omitempty"`
	Schedule  string           `json:"schedule,omitempty"`
	Manual    bool             `json:"manual,omitempty"`
	// job active deadline seconds.
	ActiveDeadlineSeconds int64 `json:"activeDeadlineSeconds,omitempty"`
}
