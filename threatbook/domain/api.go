package domain

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
	SubDomains(req *QueryRequest) (*threatbook.Response, error)
}

type api struct {
}

func (s *api) Query(req *QueryRequest) (*threatbook.Response, error) {
	res, err := threatbook.NewRest().Get(string(dict.API_V3_DOMAIN_QUERY), req)

	return res, err
}

func (s *api) AdvQuery(req *QueryRequest) (*threatbook.Response, error) {
	res, err := threatbook.NewRest().Get(string(dict.API_V3_DOMAIN_ADV_QUERY), req)

	return res, err
}

func (s *api) SubDomains(req *QueryRequest) (*threatbook.Response, error) {
	res, err := threatbook.NewRest().Get(string(dict.API_V3_DOMAIN_SUB_DOMAINS), req)

	return res, err
}
