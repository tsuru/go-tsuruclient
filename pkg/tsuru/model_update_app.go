/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type UpdateApp struct {
	// App tags.
	Tags []string `json:"tags,omitempty"`
	// App plan name.
	Plan         string       `json:"plan,omitempty"`
	Planoverride PlanOverride `json:"planoverride,omitempty"`
	// App pool name.
	Pool string `json:"pool,omitempty"`
	// App platform.
	Platform string `json:"platform,omitempty"`
	// App description.
	Description string `json:"description,omitempty"`
	// Team that owns the app.
	TeamOwner string `json:"teamOwner,omitempty"`
	// Prevent app restart.
	NoRestart bool `json:"noRestart,omitempty"`
	// Reset app image to platform base image.
	ImageReset bool         `json:"imageReset,omitempty"`
	Metadata   Metadata     `json:"metadata,omitempty"`
	Processes  []AppProcess `json:"processes,omitempty"`
}
