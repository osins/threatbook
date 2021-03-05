package ip

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
}

// API type
type api struct {
}

//  inner method
func (s *api) Query(req *QueryRequest) (*api_root.Response, error) {
	res, err := api_root.NewRest().Get(string(dict.APIV3IpQuery), req)

	return res, err
}

//  inner method
func (s *api) AdvQuery(req *QueryRequest) (*api_root.Response, error) {
	res, err := api_root.NewRest().Get(string(dict.APIV3IpAdvQuery), req)

	return res, err
}
