package domain

import "tb.humanrisk.cn/threat/tb/dict"

// QueryRequest type
type QueryRequest struct {
	Resource string           `json:"resource"`
	Exclude  dict.ExcludeType `json:"exclude"`
	Lang     dict.LangType    `json:"lang"`
}
