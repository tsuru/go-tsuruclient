/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.15
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type TeamUpdateArgs struct {
	Newname string   `json:"newname,omitempty"`
	Tags    []string `json:"tags,omitempty"`
}
