package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sdejesusp/fiber-trivia/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/welcome", handlers.Welcome)
	fact := app.Group("/fact")
	fact.Get("/", handlers.ListFacts)
	fact.Post("/", handlers.CreateFact)
	fact.Get("/:id", handlers.GetSingleFact)
	fact.Put("/:id", handlers.UpdateFact)
	fact.Delete("/:id", handlers.DeleteFact)
}
