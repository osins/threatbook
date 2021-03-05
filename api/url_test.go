package api

import (
	"fmt"
	"net/http"
	"testing"

	api_route "tb.humanrisk.cn/api/dict"
	"tb.humanrisk.cn/threat/tb/dict"
)

// TestScan method
func TestScan(t *testing.T) {
	conf := &WebAPITestConfig{
		T:           t,
		Description: "threatbook ip query",
		Method:      string(dict.RestMethodGet),
		Route:       string(api_route.RouteV1UrlScan) + "?url=" + "http://baidu.com&lang=en",
		Error:       false,
		Code:        200,
		Body:        `{"response_code":0,"verbose_msg":"成功"}`,
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

// TestReport method
func TestReport(t *testing.T) {
	conf := &WebAPITestConfig{
		T:           t,
		Description: "threatbook ip query",
		Method:      string(dict.RestMethodGet),
		Route:       string(api_route.RouteV1UrlReport) + "?url=" + "http://baidu.com&lang=en",
		Error:       false,
		Code:        200,
		Body:        `{"response_code":0,"verbose_msg":"成功"}`,
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
