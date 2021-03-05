package api

import (
	"github.com/gofiber/fiber/v2"
	"tb.humanrisk.cn/api/dict"
	tb_scene "tb.humanrisk.cn/threat/tb/scene"
)

// NewSceneRoute method
func NewSceneRoute() Routes {
	return &scene{}
}

// scene type
type scene struct {
}

// RouteInit inner method
func (s *scene) RouteInit(app *fiber.App) error {
	app.Get(string(dict.RouteV1SceneIPReputation), s.IPReputation)
	app.Get(string(dict.RouteV1SceneDNS), s.DNS)
	app.Get(string(dict.RouteV1SceneDomainContext), s.DomainContext)
	return nil
}

// IPReputation inner method
func (s *scene) IPReputation(c *fiber.Ctx) error {
	req := &tb_scene.QueryRequest{}
	err := c.QueryParser(req)
	if err != nil {
		return err
	}

	resp, err := tb_scene.New().IPReputation(req)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}

// DNS inner method
func (s *scene) DNS(c *fiber.Ctx) error {
	req := &tb_scene.QueryRequest{}
	err := c.QueryParser(req)
	if err != nil {
		return err
	}

	resp, err := tb_scene.New().DNS(req)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}

// DomainContext inner method
func (s *scene) DomainContext(c *fiber.Ctx) error {
	req := &tb_scene.QueryRequest{}
	err := c.QueryParser(req)
	if err != nil {
		return err
	}

	resp, err := tb_scene.New().DomainContext(req)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}
