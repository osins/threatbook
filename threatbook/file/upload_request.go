package file

import "tb.humanrisk.cn/threatbook/dict"

type UploadRequest struct {
	File        *dict.FileType   `json:"file"`
	SandboxType dict.SandboxType `json:"sandbox_type"`
	RunTime     int              `json:"run_time"`
}
