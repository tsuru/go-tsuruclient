/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// Auto Scale schedules struct
type AutoScaleSchedule struct {
	Name        string `json:"name,omitempty"`
	MinReplicas int32  `json:"minReplicas,omitempty"`
	Start       string `json:"start,omitempty"`
	End         string `json:"end,omitempty"`
	Timezone    string `json:"timezone,omitempty"`
}
