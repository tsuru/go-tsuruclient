/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.22
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type CertificateSetData struct {
	Cname       string `json:"cname,omitempty"`
	Certificate []byte `json:"certificate,omitempty"`
	Key         []byte `json:"key,omitempty"`
}
