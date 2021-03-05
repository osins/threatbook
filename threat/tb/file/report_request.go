package file

import (
	"tb.humanrisk.cn/threat/tb/dict"
)

// ReportRequest type
type ReportRequest struct {
	SandboxType dict.SandboxType `json:"sandbox_type"`
	Sha256      string           `json:"sha256"`
	QueryFields string           `json:"query_fields"`
}
