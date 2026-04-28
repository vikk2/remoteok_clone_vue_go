package handlers

import (
	"backend/localdata"

	"github.com/gofiber/fiber/v2"
)

func CardInfoHandler(c *fiber.Ctx) error {
	data := localdata.InitCardInfoData()
	return c.JSON(data.GetAll())
}
