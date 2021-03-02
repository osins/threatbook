package ip

import (
	"tb.humanrisk.cn/threatbook"
	"tb.humanrisk.cn/threatbook/dict"
)

func New() API {
	return &api{}
}

type API interface {
	Query(req *QueryRequest) (*threatbook.Response, error)
	AdvQuery(req *QueryRequest) (*threatbook.Response, error)
}

type api struct {
}

func (s *api) Query(req *QueryRequest) (*threatbook.Response, error) {
	res, err := threatbook.NewRest().Get(string(dict.API_V3_IP_QUERY), req)

	return res, err
}

func (s *api) AdvQuery(req *QueryRequest) (*threatbook.Response, error) {
	res, err := threatbook.NewRest().Get(string(dict.API_V3_IP_ADV_QUERY), req)

	return res, err
}
