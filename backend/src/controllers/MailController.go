package controllers

import (
	. "api/src/repositories"
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type MailController struct {
}

var mailRep MailRepository

func (uc *MailController) Send(c *fiber.Ctx) error {
	_, err := mailRep.Send(
		json.NewDecoder(bytes.NewReader(c.Body())))

	if err != nil {
		c.Status(http.StatusNotAcceptable).SendString("invalid data!")
		return nil
	}

	c.Status(http.StatusOK).SendString("Email successfully sent!")
	return nil
}
