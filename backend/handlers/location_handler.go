package handlers

import (
	"github.com/gofiber/fiber/v2"
	"backend/localdata"
)

func LocationDataHandler(c *fiber.Ctx) error {
	data := localdata.InitLocationData()
	return c.JSON(data.GetAll())
}