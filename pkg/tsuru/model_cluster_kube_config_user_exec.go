/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.19
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type ClusterKubeConfigUserExec struct {
	Args       []string                       `json:"args,omitempty"`
	ApiVersion string                         `json:"apiVersion,omitempty"`
	Env        []ClusterKubeConfigUserExecEnv `json:"env,omitempty"`
	Command    string                         `json:"command,omitempty"`
}
