/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.15
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type SetRoutableArgs struct {
	Version    string `json:"version,omitempty"`
	IsRoutable bool   `json:"isRoutable,omitempty"`
}
