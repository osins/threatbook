package file

import "tb.humanrisk.cn/threat/tb/dict"

// UploadRequest type
type UploadRequest struct {
	File        *dict.FileType   `json:"file"`
	SandboxType dict.SandboxType `json:"sandbox_type"`
	RunTime     int              `json:"run_time"`
}
