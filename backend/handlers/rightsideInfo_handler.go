package handlers

import (
	"github.com/gofiber/fiber/v2"
	"backend/localdata"
)

func RigthSideInfoHandler(c *fiber.Ctx) error{
	data := localdata.InitRightSideInfoData()
	return c.JSON(data.GetAll())
}