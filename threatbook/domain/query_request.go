package domain

import "tb.humanrisk.cn/threatbook/dict"

type QueryRequest struct {
	Resource string           `json:"resource"`
	Exclude  dict.ExcludeType `json:"exclude"`
	Lang     dict.LangType    `json:"lang"`
}
