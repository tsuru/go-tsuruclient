/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type VolumeBindData struct {
	App        string `json:"app,omitempty"`
	Mountpoint string `json:"mountpoint,omitempty"`
	Norestart  bool   `json:"norestart,omitempty"`
	Readonly   bool   `json:"readonly,omitempty"`
}
