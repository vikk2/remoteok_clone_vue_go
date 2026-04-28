package routes

import (
	"backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func BtnDataRoutes(app *fiber.App) {
	app.Get("/api/btn-data", handlers.BtnDataHandler)
}

func BenefitDataRoutes(app *fiber.App) {
	app.Get("/api/benefit-data", handlers.BenefitHandler)
}

func PriceRangeDataRoutes(app *fiber.App) {
	app.Get("/api/price-range-data", handlers.PriceRangeHandler)
}

func CardInfoRoutes(app *fiber.App) {
	app.Get("/api/card-info", handlers.CardInfoHandler)
}

func LocationDataRoutes(app *fiber.App) {
	app.Get("/api/location-data", handlers.LocationDataHandler)
}

func LogoDropdownDataRoutes(app *fiber.App) {
	app.Get("/api/logo-dropdown-data", handlers.LogoDropDownHandler)
}

func TabLabelDataRoutes(app *fiber.App) {
	app.Get("/api/tab-label-data", handlers.TabLabelHandler)
}

func SortByRoutes(app *fiber.App){
	app.Get("/api/sortby-data", handlers.SortByHandler)
}

func SearchTabRoutes(app *fiber.App){
	app.Get("/api/search-tab-data", handlers.SearchTabHandler)
}

func PostJobLogoRoutes(app *fiber.App){
	app.Get("/api/post-job-logo", handlers.PostJobLogoCompanyHandler)
}

func RightSideInfoRoutes(app *fiber.App){
	app.Get("/api/right-side-info", handlers.RigthSideInfoHandler)
}

func FeedbackRoutes(app *fiber.App){
	app.Get("/api/feedback-data", handlers.UserFeedHandler)
}