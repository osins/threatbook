package api

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	api_route "tb.humanrisk.cn/api/dict"
	"tb.humanrisk.cn/threat/tb/dict"
)

// TestDomainQuery method
func TestDomainQuery(t *testing.T) {
	conf := &WebAPITestConfig{
		T:           t,
		Description: "domain query",
		Method:      string(dict.RestMethodGet),
		Route:       fmt.Sprintf("%s?resource=%s&lang=%s", api_route.RouteV1DomainQuery, "baidu.com", "en"),
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

	assert.Nilf(t, err, conf.Description)

	wat.Run(req)
}

// TestDomainAdvQuery method
func TestDomainAdvQuery(t *testing.T) {
	conf := &WebAPITestConfig{
		T:           t,
		Description: "domain adv query",
		Method:      string(dict.RestMethodGet),
		Route:       fmt.Sprintf("%s?resource=%s&lang=%s", api_route.RouteV1DomainAdvQuery, "baidu.com", "en"),
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

	assert.Nilf(t, err, conf.Description)

	wat.Run(req)
}

// TestDomainSubDomains method
func TestDomainSubDomains(t *testing.T) {
	conf := &WebAPITestConfig{
		T:           t,
		Description: "domain sub query",
		Method:      string(dict.RestMethodGet),
		Route:       fmt.Sprintf("%s?resource=%s&lang=%s", api_route.RouteV1DomainSubDomains, "baidu.com", "en"),
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

	assert.Nilf(t, err, conf.Description)

	wat.Run(req)
}
