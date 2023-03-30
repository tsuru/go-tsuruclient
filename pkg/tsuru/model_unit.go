/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.14
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type Unit struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Appname     string `json:"appname,omitempty"`
	Processname string `json:"processname,omitempty"`
	Type        string `json:"type,omitempty"`
	Ip          string `json:"ip,omitempty"`
	Status      string `json:"status,omitempty"`
	Version     int32  `json:"version,omitempty"`
	Routable    *bool  `json:"routable,omitempty"`
	Ready       *bool  `json:"ready,omitempty"`
	Restarts    *int   `json:"restarts,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	Address     Url    `json:"address,omitempty"`
}
