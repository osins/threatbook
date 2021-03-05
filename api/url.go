package api

import (
	"github.com/gofiber/fiber/v2"
	"tb.humanrisk.cn/api/dict"
	tb_url "tb.humanrisk.cn/threat/tb/url"
)

// NewURLRoute method
func NewURLRoute() Routes {
	return &url{}
}

// url type
type url struct {
}

// RouteInit inner method
func (s *url) RouteInit(app *fiber.App) error {
	app.Get(string(dict.RouteV1UrlScan), s.Scan)
	app.Get(string(dict.RouteV1UrlReport), s.Report)
	return nil
}

// Scan inner method
func (s *url) Scan(c *fiber.Ctx) error {
	req := &tb_url.QueryRequest{}
	err := c.QueryParser(req)
	if err != nil {
		return err
	}

	resp, err := tb_url.New().Scan(req)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}

// Report inner method
func (s *url) Report(c *fiber.Ctx) error {
	req := &tb_url.QueryRequest{}
	err := c.QueryParser(req)
	if err != nil {
		return err
	}

	resp, err := tb_url.New().Report(req)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}
