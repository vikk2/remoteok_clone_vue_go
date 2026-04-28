package handlers

import (
	"github.com/gofiber/fiber/v2"

	"backend/localdata"
)

func BtnDataHandler(c *fiber.Ctx) error {
	data := localdata.InitBtnData()
	return c.JSON(data.GetAll())
}