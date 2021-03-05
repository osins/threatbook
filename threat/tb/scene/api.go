package scene

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
	IPReputation(req *QueryRequest) (*api_root.Response, error)
	DNS(req *QueryRequest) (*api_root.Response, error)
	DomainContext(req *QueryRequest) (*api_root.Response, error)
}

// API type
type api struct {
}

//  inner method
func (s *api) IPReputation(req *QueryRequest) (*api_root.Response, error) {
	res, err := api_root.NewRest().Get(string(dict.APIV3ScenIPReputation), req)

	return res, err
}

//  inner method
func (s *api) DNS(req *QueryRequest) (*api_root.Response, error) {
	res, err := api_root.NewRest().Get(string(dict.APIV3SceneDNS), req)

	return res, err
}

//  inner method
func (s *api) DomainContext(req *QueryRequest) (*api_root.Response, error) {
	res, err := api_root.NewRest().Get(string(dict.APIV3SceneDomainContext), req)

	return res, err
}
