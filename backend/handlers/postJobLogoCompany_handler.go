package handlers

import (
	"github.com/gofiber/fiber/v2"
	"backend/localdata"
)

func PostJobLogoCompanyHandler(c *fiber.Ctx) error{
	data := localdata.InitPostJobLogoData()
	return c.JSON(data.GetAll())
}