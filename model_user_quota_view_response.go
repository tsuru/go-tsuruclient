/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru
// UserQuotaViewResponse Response returned by User Quota View.
type UserQuotaViewResponse struct {
	Inuse int32 `json:"inuse,omitempty"`
	Limit int32 `json:"limit,omitempty"`
}
