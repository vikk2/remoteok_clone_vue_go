package main

import (
	"backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173/, http://127.0.0.1:5173/",
	}))

	routes.BtnDataRoutes(app)
	routes.BenefitDataRoutes(app)
	routes.PriceRangeDataRoutes(app)
	routes.CardInfoRoutes(app)
	routes.LocationDataRoutes(app)
	routes.TabLabelDataRoutes(app)
	routes.SortByRoutes(app)
	routes.SearchTabRoutes(app)
	routes.PostJobLogoRoutes(app)
	routes.LogoDropdownDataRoutes(app)
	routes.RightSideInfoRoutes(app)
	routes.FeedbackRoutes(app)

	app.Listen(":3000")

}
