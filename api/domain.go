package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"tb.humanrisk.cn/api/dict"
	api_root "tb.humanrisk.cn/threat/tb"
	td "tb.humanrisk.cn/threat/tb/domain"
)

// NewDomainRoute method
func NewDomainRoute() Routes {
	return &domain{}
}

// domain type
type domain struct {
}

// RouteInit inner method
func (s *domain) RouteInit(app *fiber.App) error {
	app.Get(string(dict.RouteV1DomainQuery), s.Query)
	app.Get(string(dict.RouteV1DomainAdvQuery), s.AdvQuery)
	app.Get(string(dict.RouteV1DomainSubDomains), s.SubDomains)
	return nil
}

// Query inner method
func (s *domain) Query(c *fiber.Ctx) error {
	fmt.Println("request body: ", c.Context().QueryArgs())

	req := &td.QueryRequest{}
	err := c.QueryParser(req)
	if err != nil {
		return err
	}

	var resp *api_root.Response
	resp, err = td.New().Query(req)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}

// AdvQuery inner method
func (s *domain) AdvQuery(c *fiber.Ctx) error {
	fmt.Println("request body: ", c.Context().QueryArgs())

	req := &td.QueryRequest{}
	err := c.QueryParser(req)
	if err != nil {
		return err
	}

	var resp *api_root.Response
	resp, err = td.New().AdvQuery(req)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}

// SubDomains inner method
func (s *domain) SubDomains(c *fiber.Ctx) error {
	fmt.Println("request body: ", c.Context().QueryArgs())

	req := &td.QueryRequest{}
	err := c.QueryParser(req)
	if err != nil {
		return err
	}

	var resp *api_root.Response
	resp, err = td.New().SubDomains(req)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}
