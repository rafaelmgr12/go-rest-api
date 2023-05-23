package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	// create a fiber application
	var app *fiber.App = fiber.New()

	// add a request handler
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// start the application on port 3000
	app.Listen(":3000")
}
