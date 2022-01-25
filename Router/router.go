package router

import (
	h "locationhistory/handler"

	"github.com/gofiber/fiber/v2"
)

//SetupRoutes func initializes the server routes
func SetupRoutes(app *fiber.App) {
	//group location routes
	location := app.Group("location")

	location.Post("/:orderId/now", h.AppendLocation)

	location.Get("/:orderId", h.GetHistory)

	location.Delete("/:orderId", h.DeleteHistory)

	//handles all undefined routes
	app.Use(func(c *fiber.Ctx) error {
		return h.HomePage(c)
	})
}
