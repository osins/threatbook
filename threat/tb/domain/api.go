package domain

import (
	api_root "tb.humanrisk.cn/threat/tb"
	"tb.humanrisk.cn/threat/tb/dict"
)

// New method
func New() API {
	return &api{}
}

// API interface
type API interface {
	Query(req *QueryRequest) (*api_root.Response, error)
	AdvQuery(req *QueryRequest) (*api_root.Response, error)
	SubDomains(req *QueryRequest) (*api_root.Response, error)
}

// API type
type api struct {
}

//  inner method
func (s *api) Query(req *QueryRequest) (*api_root.Response, error) {
	res, err := api_root.NewRest().Get(string(dict.APIV3DomainQuery), req)

	return res, err
}

//  inner method
func (s *api) AdvQuery(req *QueryRequest) (*api_root.Response, error) {
	res, err := api_root.NewRest().Get(string(dict.APIV3DomainAdvQuery), req)

	return res, err
}

//  inner method
func (s *api) SubDomains(req *QueryRequest) (*api_root.Response, error) {
	res, err := api_root.NewRest().Get(string(dict.APIV3DomainSubDomains), req)

	return res, err
}
