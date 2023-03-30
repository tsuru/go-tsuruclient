/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.14
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type JobSpec struct {
	Container InputJobContainer `json:"container,omitempty"`
	Schedule  string            `json:"schedule,omitempty"`
}