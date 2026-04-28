package handlers

import (
	"github.com/gofiber/fiber/v2"
	"backend/localdata"
)

func LogoDropDownHandler(c *fiber.Ctx) error{
	data := localdata.InitLogoDropDownData()
	return c.JSON(data.GetAll())
}