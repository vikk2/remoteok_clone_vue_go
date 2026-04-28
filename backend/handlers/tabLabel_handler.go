package handlers

import (
	"github.com/gofiber/fiber/v2"
	"backend/localdata"
)

func TabLabelHandler(c *fiber.Ctx) error{
	data := localdata.InitLabelData()
	return c.JSON(data.GetAll())
}
