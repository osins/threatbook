package scene

import (
	"tb.humanrisk.cn/threatbook"
	"tb.humanrisk.cn/threatbook/dict"
)

func New() API {
	return &api{}
}

type API interface {
	IpReputation(req *QueryRequest) (*threatbook.Response, error)
	DNS(req *QueryRequest) (*threatbook.Response, error)
	DomainContext(req *QueryRequest) (*threatbook.Response, error)
}

type api struct {
}

func (s *api) IpReputation(req *QueryRequest) (*threatbook.Response, error) {
	res, err := threatbook.NewRest().Get(string(dict.API_V3_SCEN_IP_REPUTATION), req)

	return res, err
}

func (s *api) DNS(req *QueryRequest) (*threatbook.Response, error) {
	res, err := threatbook.NewRest().Get(string(dict.API_V3_SCENE_DNS), req)

	return res, err
}

func (s *api) DomainContext(req *QueryRequest) (*threatbook.Response, error) {
	res, err := threatbook.NewRest().Get(string(dict.API_V3_SCENE_DOMAIN_CONTEXT), req)

	return res, err
}
