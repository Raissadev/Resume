package middlewares

import (
	r "api/src/repositories"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

type SecMiddleware struct{}

var rep r.IpRepository

func (sm *SecMiddleware) VToken(ctx *fiber.Ctx) (err error) {
	tk := ctx.Get("X-Session-Token")
	tk_sum := os.Getenv("KEY_VALUE")

	if tk != tk_sum {
		ctx.Status(http.StatusNotAcceptable).JSON(fiber.Map{
			"message": "Invalid token.",
		})
		return
	}

	ctx.Next()
	return
}

func (sm *SecMiddleware) Ip(c *fiber.Ctx) error {
	userAgent := c.Get("User-Agent")
	clientIP := c.IP()
	c.Locals("clientIP", clientIP)
	rep.Create(clientIP, userAgent)
	return c.Next()
}
