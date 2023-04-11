package routes

import (
	controllers "api/src/controllers"
	. "api/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	logger_fiber "github.com/gofiber/fiber/v2/middleware/logger"
)

var env = Lenv.New()
var chc controllers.HalfCheckController
var cml controllers.MailController

type Routes struct {
	Server *fiber.App
	RestV1 fiber.Router
}

func New() *Routes {
	app := &Routes{
		Server: fiber.New(),
	}

	app.Server.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	app.Server.Use(compress.New())
	app.Server.Use(logger_fiber.New())

	app.Server.Static("/", "/public", fiber.Static{
		Compress: true,
		Browse:   true,
	})

	app.RestV1 = app.Server.Group("/api/v1")

	return app
}

func (app *Routes) NewPublicAPI() *fiber.App {
	app.RestV1.Get("/", chc.New)
	app.RestV1.Post("/mail", cml.Send)
	return app.Server
}

func (app *Routes) NewWeb() {
	app.Server.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("/public/index.html")
	})
}
