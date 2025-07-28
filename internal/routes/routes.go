package routes

import (
	"github.com/crazyuploader/GoAPI/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Root endpoint
	app.Get("/", handlers.RootHandler)

	// Health check endpoint
	app.Get("/health", handlers.HealthCheck)
}
