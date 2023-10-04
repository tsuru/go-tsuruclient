/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.18
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type ServiceBroker struct {
	Name   string              `json:"Name,omitempty"`
	URL    string              `json:"URL,omitempty"`
	Config ServiceBrokerConfig `json:"Config,omitempty"`
}
