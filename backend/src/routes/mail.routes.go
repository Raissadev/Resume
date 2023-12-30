package routes

import (
	c "api/src/controllers"
	m "api/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func (app *Routes) Mail() *fiber.App {
	var chc c.HalfCheckController
	var cml c.MailController
	var sec m.SecMiddleware

	// Log ip
	app.RestV1.Use(sec.Ip)

	// Half check
	app.RestV1.Get("/", chc.New)

	// Send mail
	app.RestV1.Post("/mail", cml.Send)

	return app.Server
}
