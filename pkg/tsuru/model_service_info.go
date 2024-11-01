/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type ServiceInfo struct {
	Id          string                 `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Pool        string                 `json:"pool,omitempty"`
	Teams       []string               `json:"teams,omitempty"`
	Planname    string                 `json:"planname,omitempty"`
	Apps        []string               `json:"apps,omitempty"`
	Jobs        []string               `json:"jobs,omitempty"`
	Servicename string                 `json:"servicename,omitempty"`
	Info        map[string]interface{} `json:"info,omitempty"`
	Teamowner   string                 `json:"teamowner,omitempty"`
}
