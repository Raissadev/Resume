package controllers

import (
	. "api/src/repositories"
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
}

var userRep UserRepository

func (uc *UserController) List(c *fiber.Ctx) error {
	resp := (&UserRepository{}).All()

	c.Status(http.StatusOK).JSON(resp)
	return nil
}

func (uc *UserController) Show(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		c.Status(http.StatusNotFound).SendString("error!")
		return nil
	}

	resp, err := userRep.Get(id)
	if err != nil {
		c.Status(http.StatusNotFound).SendString("not found!")
		return nil
	}

	c.Status(http.StatusOK).JSON(resp)
	return nil
}

func (uc *UserController) Store(c *fiber.Ctx) error {
	_, err := userRep.Create(json.NewDecoder(bytes.NewReader(c.Body())))

	if err != nil {
		return c.Status(http.StatusNotAcceptable).SendString("invalid data")
	}

	return c.SendString("successfully created!")
}

func (uc *UserController) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		c.Status(http.StatusNotFound).SendString("error!")
		return nil
	}

	resp, err := userRep.Update(id, json.NewDecoder(bytes.NewReader(c.Body())))
	if err != nil {
		c.Status(http.StatusNotFound).SendString(err.Error())
		return nil
	}

	c.Status(http.StatusOK).JSON(resp)
	return nil
}

func (uc *UserController) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(http.StatusNotFound).SendString("not found")
	}

	resp, err := userRep.Delete(id)

	if err != nil || !resp {
		return c.Status(http.StatusNotAcceptable).SendString(err.Error())
	}

	return c.SendString("successfully deleted!")
}
