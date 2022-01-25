package main

import (
	"locationhistory/config"
	r "locationhistory/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	r.SetupRoutes(app)

	port := config.Config("HISTORY_SERVER_LISTEN_ADDR")

	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
