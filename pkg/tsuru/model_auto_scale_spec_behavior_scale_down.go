/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type AutoScaleSpecBehaviorScaleDown struct {
	// Time in seconds for stabilization before scaling down
	StabilizationWindow int32 `json:"stabilizationWindow,omitempty"`
	// Minimum number of units to scale down
	UnitsPolicyValue int32 `json:"unitsPolicyValue,omitempty"`
	// Percentage threshold for scaling down
	PercentagePolicyValue int32 `json:"percentagePolicyValue,omitempty"`
}