package handlers

import (
	"github.com/gofiber/fiber/v2"
	"backend/localdata"
)

func SortByHandler(c *fiber.Ctx) error{
	data := localdata.InitSortByData()
	return c.JSON(data.GetAll())
}