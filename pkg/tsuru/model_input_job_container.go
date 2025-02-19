/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type InputJobContainer struct {
	Image   string   `json:"image,omitempty"`
	Envs    []EnvVar `json:"envs,omitempty"`
	Command []string `json:"command,omitempty"`
}
