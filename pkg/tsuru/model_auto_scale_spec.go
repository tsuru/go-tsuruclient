/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// Units Auto Scale spec
type AutoScaleSpec struct {
	Process    string                `json:"process,omitempty"`
	MinUnits   int32                 `json:"minUnits,omitempty"`
	MaxUnits   int32                 `json:"maxUnits,omitempty"`
	AverageCPU string                `json:"averageCPU,omitempty"`
	Schedules  []AutoScaleSchedule   `json:"schedules,omitempty"`
	Prometheus []AutoScalePrometheus `json:"prometheus,omitempty"`
	Version    int32                 `json:"version,omitempty"`
	Behavior   AutoScaleSpecBehavior `json:"behavior,omitempty"`
}
