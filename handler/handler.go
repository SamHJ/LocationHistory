package handler

import (
	m "locationhistory/model"

	"github.com/gofiber/fiber/v2"
)

//HomePage func returns the static page of the server
func HomePage(c *fiber.Ctx) error {
	return c.Status(200).SendFile("./public/index.html")
}

func AppendLocation(c *fiber.Ctx) error {
	orderId := c.Params("orderId")
	history := m.History{
		Lat: 12.34,
		Lng: 56.78,
	}

	history.StoreHistory(orderId)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "History stored successfully",
		"data":    history,
	})
}
