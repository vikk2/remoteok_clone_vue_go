package handlers

import (
	"github.com/gofiber/fiber/v2"

	"backend/localdata"
)

func SearchTabHandler(c *fiber.Ctx) error{
	data := localdata.InitSearchTabData()
	return c.JSON(data.GetAll())
}