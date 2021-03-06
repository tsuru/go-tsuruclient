/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type SshKeyAddData struct {
	Force   bool   `json:"force,omitempty"`
	Key     string `json:"key,omitempty"`
	Keyname string `json:"keyname,omitempty"`
}
