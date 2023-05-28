package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelmgr12/go-rest-api/internal/database"
	"github.com/rafaelmgr12/go-rest-api/internal/routes"
	"github.com/rafaelmgr12/go-rest-api/internal/utils"
)

const DEFAULT_PORT = "3000"

func NewFiberApp() *fiber.App {
	var app *fiber.App = fiber.New()

	routes.SetupRoutes(app)

	return app
}

func main() {
	var app *fiber.App = NewFiberApp()

	database.InitDatabase(utils.GetValue("DB_NAME"))

	var PORT string = os.Getenv("PORT")

	if PORT == "" {
		PORT = DEFAULT_PORT
	}

	app.Listen(fmt.Sprintf(":%s", PORT))
}
