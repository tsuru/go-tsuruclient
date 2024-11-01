/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type AppId struct {
	App        string `json:"App,omitempty"`
	MountPoint string `json:"MountPoint,omitempty"`
	Volume     string `json:"Volume,omitempty"`
}
