package scene

import "tb.humanrisk.cn/threat/tb/dict"

// ReputationRequest type
type ReputationRequest struct {
	Resource string        `json:"resource"`
	Lang     dict.LangType `json:"lang"`
}
