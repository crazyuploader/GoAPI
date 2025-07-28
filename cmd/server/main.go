package main

import (
	"log"

	"github.com/crazyuploader/GoAPI/internal/config"
	"github.com/crazyuploader/GoAPI/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load Environment Variables from .env file. Falling back to Environment Variables.")
	}

	// Load configuration
	cfg := config.Load()

	log.Printf("Starting API Server in %v mode.", cfg.Environment)

	// Initialize Fiber app with optimized settings for performance
	app := fiber.New(fiber.Config{
		AppName:      cfg.AppName,
		ServerHeader: "Fiber",
		Prefork:      false,
	})

	// Enable middleware(s)
	app.Use(logger.New())

	// Setup routes
	routes.SetupRoutes(app)

	// Set API port
	port := cfg.Port
	if port == "" {
		port = "3100"
	}

	// Start API Server
	log.Printf("🚀 Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
