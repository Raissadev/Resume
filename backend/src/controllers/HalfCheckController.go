package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type HalfCheckController struct {
}

var halfService HalfCheckController

func (hf *HalfCheckController) New(c *fiber.Ctx) (retErr error) {
	c.Status(http.StatusOK).SendString("42")
	return
}
