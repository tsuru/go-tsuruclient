/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.15
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type RecommendedResources struct {
	Process         string                                `json:"process,omitempty"`
	Recommendations []RecommendedResourcesRecommendations `json:"recommendations,omitempty"`
}