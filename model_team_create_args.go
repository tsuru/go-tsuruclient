/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru
// TeamCreateArgs struct for TeamCreateArgs
type TeamCreateArgs struct {
	Name string `json:"name,omitempty"`
	Tags []string `json:"tags,omitempty"`
}
