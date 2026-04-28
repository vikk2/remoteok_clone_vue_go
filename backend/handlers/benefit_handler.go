package handlers

import (
	"github.com/gofiber/fiber/v2"
	"backend/localdata"
)

func BenefitHandler(c *fiber.Ctx) error {
	data := localdata.InitBenefitData()
	return c.JSON(data.GetAll())
}