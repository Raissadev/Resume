package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type HalfCheckController struct{}

func (hf *HalfCheckController) New(c *fiber.Ctx) (err error) {
	// absolutely nothing
	c.Status(http.StatusOK).SendString("42")
	return
}
