/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.20
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type ServiceInstanceUnbind struct {
	NoRestart bool `json:"noRestart,omitempty"`
	Force     bool `json:"force,omitempty"`
}
