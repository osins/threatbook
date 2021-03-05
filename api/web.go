package api

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"tb.humanrisk.cn/docs"
)

// Boot 启动web服务
// Boot method
func Boot() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	SwaggerInit(app)
	RouteInit(app)

	app.Listen(":8089")

	return app
}

// SwaggerInit 初始化API文档 run command: swag init
// SwaggerInit method
func SwaggerInit(app *fiber.App) {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "威胁情报 API"
	docs.SwaggerInfo.Description = "这是威胁情报相关接口服务"
	docs.SwaggerInfo.Version = "v0.0.1"
	docs.SwaggerInfo.Host = "tb.humanrisk.cn"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	app.Get("/swagger/*", swagger.Handler) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://localhost:8089/swagger/doc.json",
		DeepLinking: false,
	}))
}

// RouteInit 初始化路由
// RouteInit method
func RouteInit(app *fiber.App) {
	NewIPRoute().RouteInit(app)
	NewDomainRoute().RouteInit(app)
	NewFileRoute().RouteInit(app)
	NewSceneRoute().RouteInit(app)
	NewURLRoute().RouteInit(app)
}
