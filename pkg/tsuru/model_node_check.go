/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.18
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

import (
	"time"
)

type NodeCheck struct {
	Time   time.Time         `json:"time,omitempty"`
	Checks []NodeCheckResult `json:"checks,omitempty"`
}
