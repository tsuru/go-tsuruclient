/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.22
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// Auto Scale prometheus struct
type AutoScalePrometheus struct {
	Name                string  `json:"name,omitempty"`
	Query               string  `json:"query,omitempty"`
	Threshold           float64 `json:"threshold,omitempty"`
	ActivationThreshold float64 `json:"activationThreshold,omitempty"`
	PrometheusAddress   string  `json:"prometheusAddress,omitempty"`
}
