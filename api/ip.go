package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"tb.humanrisk.cn/api/dict"
	api_root "tb.humanrisk.cn/threat/tb"
	tip "tb.humanrisk.cn/threat/tb/ip"
)

// NewIPRoute method
func NewIPRoute() Routes {
	return &ip{}
}

// ip type
type ip struct {
}

// RouteInit inner method
func (s *ip) RouteInit(app *fiber.App) error {
	app.Get(string(dict.RouteV1IpQuery), s.Query)
	app.Get(string(dict.RouteV1IpAdvQuery), s.AdvQuery)
	return nil
}

// IP 分析
// @Summary 可针对业务日志，以及从防火墙、WAF等安防设备中获取的外部IP，进行分析。
// @Description 获取IP相关地理位置、ASN信息，综合判定威胁类型如:远控（C2）、傀儡机（Zombie）、失陷主机（Compromised）、扫描（Scanner）、钓鱼（Phishing）等，相关攻击团伙或安全事件标签，原始情报，相关样本信息等。
// @ID ip-query
// @Accept  json
// @Produce  json
// @Param resource query string false "resource string"
// @Param exclude query string false "exclude string"
// @Param lang query string false "lang string"
// @Success 200 {object} api_root.Response
// @Failure 400 {string} string	"error"
// @Failure 404 {string} string	"error"
// @Failure 500 {string} string	"error"
// @Router /ip/query [get]
// Query inner method
func (s *ip) Query(c *fiber.Ctx) error {
	fmt.Println("request body: ", c.Context().QueryArgs())

	req := &tip.QueryRequest{}
	err := c.QueryParser(req)
	if err != nil {
		return err
	}

	var resp *api_root.Response
	resp, err = tip.New().Query(req)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}

// AdvQuery inner method
func (s *ip) AdvQuery(c *fiber.Ctx) error {
	req := &tip.QueryRequest{}
	err := c.QueryParser(req)
	if err != nil {
		return err
	}

	var resp *api_root.Response
	resp, err = tip.New().AdvQuery(req)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}
