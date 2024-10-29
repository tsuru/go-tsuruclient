/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type Webhook struct {
	Name        string              `json:"name,omitempty"`
	Description string              `json:"description,omitempty"`
	TeamOwner   string              `json:"team_owner,omitempty"`
	EventFilter WebhookEventFilter  `json:"event_filter,omitempty"`
	Url         string              `json:"url,omitempty"`
	ProxyUrl    string              `json:"proxy_url,omitempty"`
	Headers     map[string][]string `json:"headers,omitempty"`
	Method      string              `json:"method,omitempty"`
	Body        string              `json:"body,omitempty"`
	Insecure    bool                `json:"insecure,omitempty"`
}
