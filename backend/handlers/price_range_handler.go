package handlers

import (
	"github.com/gofiber/fiber/v2"
	"backend/localdata"
)

func PriceRangeHandler(c *fiber.Ctx) error{
	data := localdata.InitPriceRangeData()
	return c.JSON(data.GetAll())
}