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
	Scan(req *QueryRequest) (*api_root.Response, error)
	Report(req *QueryRequest) (*api_root.Response, error)
}

// API type
type api struct {
}

//  inner method
func (s *api) Scan(req *QueryRequest) (*api_root.Response, error) {
	res, err := api_root.NewRest().Get(string(dict.APIV3UrlScan), req)

	return res, err
}

//  inner method
func (s *api) Report(req *QueryRequest) (*api_root.Response, error) {
	res, err := api_root.NewRest().Get(string(dict.APIV3UrlReport), req)

	return res, err
}
