package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafaelmgr12/go-rest-api/internal/routes"
)

func main() {

	// create a fiber application
	var app *fiber.App = fiber.New()

	// add a request handler
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.SetupRoutes(app)

	// start the application on port 3000
	app.Listen(":3000")
}
