/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.19
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type ServiceBrokerConfig struct {
	Context                map[string]string             `json:"Context,omitempty"`
	Insecure               bool                          `json:"Insecure,omitempty"`
	CacheExpirationSeconds int32                         `json:"CacheExpirationSeconds,omitempty"`
	AuthConfig             ServiceBrokerConfigAuthConfig `json:"AuthConfig,omitempty"`
}
