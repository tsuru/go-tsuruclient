/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.15
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// Assign role to token arguments.
type AssignTokenArgs struct {
	TokenId string `json:"token_id,omitempty"`
	Context string `json:"context,omitempty"`
}
