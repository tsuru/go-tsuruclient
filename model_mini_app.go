/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru
// MiniApp List containing minimal information about apps.
type MiniApp struct {
	Name string `json:"name,omitempty"`
	Pool string `json:"pool,omitempty"`
	TeamOwner string `json:"teamOwner,omitempty"`
	Plan Plan `json:"plan,omitempty"`
	Units []Unit `json:"units,omitempty"`
	Cname []string `json:"cname,omitempty"`
	Ip string `json:"ip,omitempty"`
	Routers []Router `json:"routers,omitempty"`
	Lock Lock `json:"lock,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Error string `json:"error,omitempty"`
}
