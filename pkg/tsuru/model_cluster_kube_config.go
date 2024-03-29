/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.20
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type ClusterKubeConfig struct {
	Cluster ClusterKubeConfigCluster `json:"cluster,omitempty"`
	User    ClusterKubeConfigUser    `json:"user,omitempty"`
}
