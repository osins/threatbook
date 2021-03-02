package scene

import "tb.humanrisk.cn/threatbook/dict"

type ReputationRequest struct {
	Resource string        `json:"resource"`
	Lang     dict.LangType `json:"lang"`
}
