package controllers

import (
	"api/src/config/logger"
	. "api/src/repositories"
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type MailController struct{}

var rep MailRepository

func (uc *MailController) Send(c *fiber.Ctx) error {
	_, err := rep.Send(
		json.NewDecoder(bytes.NewReader(c.Body())))

	if err != nil {
		log.Println(err)
		logger.Error("Err mail", err)
		c.Status(http.StatusNotAcceptable).JSON(fiber.Map{
			"message": "Invalid data",
		})
		return nil
	}

	c.Status(http.StatusOK).SendString("Email successfully sent!")
	return nil
}
