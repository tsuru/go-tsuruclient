/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.19
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type ClusterHelp struct {
	ProvisionerHelp string            `json:"provisioner_help,omitempty"`
	CustomDataHelp  map[string]string `json:"custom_data_help,omitempty"`
	CreateDataHelp  map[string]string `json:"create_data_help,omitempty"`
}
