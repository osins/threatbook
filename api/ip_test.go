package api

import (
	"fmt"
	"net/http"
	"testing"

	api_route "tb.humanrisk.cn/api/dict"
	"tb.humanrisk.cn/threat/tb/dict"
)

// TestIpQuery method
func TestIpQuery(t *testing.T) {
	conf := &WebAPITestConfig{
		T:           t,
		Description: "threatbook ip query",
		Method:      string(dict.RestMethodGet),
		Route:       string(api_route.RouteV1IpQuery) + "?resource=" + "158.247.226.121&lang=en",
		Error:       false,
		Code:        200,
		Body:        `{"response_code":0,"verbose_msg":"OK"}`,
	}

	wat := NewWebAPITest(conf)

	req, err := http.NewRequest(
		conf.Method,
		conf.Route,
		nil,
	)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	wat.Run(req)
}

// TestIpAdvQuery method
func TestIpAdvQuery(t *testing.T) {
	conf := &WebAPITestConfig{
		T:           t,
		Description: "threatbook adv query",
		Method:      string(dict.RestMethodGet),
		Route:       string(api_route.RouteV1IpAdvQuery) + "?resource=" + "158.247.226.121&lang=en&exclude=" + string(dict.ExcludeIntelligences),
		Error:       false,
		Code:        200,
		Body:        `{"response_code":0,"verbose_msg":"No Access to API Method"}`,
	}

	wat := NewWebAPITest(conf)

	req, err := http.NewRequest(
		conf.Method,
		conf.Route,
		nil,
	)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	wat.Run(req)
}
