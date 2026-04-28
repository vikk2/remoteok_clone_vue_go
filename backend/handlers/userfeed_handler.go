package handlers

import (
	"github.com/gofiber/fiber/v2"
	"backend/localdata"
)

func UserFeedHandler(c *fiber.Ctx) error {
	data := localdata.InitFeedbackData()
	return c.JSON(data.GetAll())
}