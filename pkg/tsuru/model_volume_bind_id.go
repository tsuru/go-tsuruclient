/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type VolumeBindId struct {
	// App the volume is bound to.
	App string `json:"app,omitempty"`
	// Volume mountpoint.
	Mountpoint string `json:"mountpoint,omitempty"`
	// Volume name.
	Volume string `json:"volume,omitempty"`
}
