/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.22
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// Arguments for creating a new team token.
type TeamTokenCreateArgs struct {
	TokenId     string `json:"token_id,omitempty"`
	Description string `json:"description,omitempty"`
	// Expire time in seconds.
	ExpiresIn int64  `json:"expires_in,omitempty"`
	Team      string `json:"team,omitempty"`
}
