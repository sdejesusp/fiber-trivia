package main

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/sdejesusp/fiber-trivia/database"
)

func main() {
	// Connection with the Postgres DB
	database.ConnectDb()

	// Fiber instance
	app := fiber.New()

	// Basic configuration for the API documentation
	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.yaml",
		Path:     "swagger",
		Title:    "Fiber Trivia API Docs",
		CacheAge: 1,
	}

	// Swagger middleware
	app.Use(swagger.New(cfg))

	// API routes
	setupRoutes(app)

	app.Get("/welcome", func(c *fiber.Ctx) error {
		return c.SendString("Hello, from a Docker Fiber")
	})

	app.Listen(":3000")
}
