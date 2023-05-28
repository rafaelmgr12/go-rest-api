package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafaelmgr12/go-rest-api/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
	// register the authentication handlers
	app.Post("/api/v1/signup", handlers.Signup)
	app.Post("/api/v1/login", handlers.Login)

	app.Get("/api/v1/items", handlers.GetAllItems)
	app.Get("/api/v1/items/:id", handlers.GetItemByID)
	app.Post("/api/v1/items", handlers.CreateItem)
	app.Put("/api/v1/items/:id", handlers.UpdateItem)
	app.Delete("/api/v1/items/:id", handlers.DeleteItem)
}
