package api

import "github.com/gofiber/fiber/v2"

// Routes interface
type Routes interface {
	RouteInit(app *fiber.App) error
}
