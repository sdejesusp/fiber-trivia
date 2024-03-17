package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sdejesusp/fiber-trivia/handlers"
)

func setupRoutes(app *fiber.App) {
	fact := app.Group("/fact")
	fact.Get("/", handlers.ListFacts)
	fact.Post("/", handlers.CreateFact)
	fact.Get("/:id", handlers.GetSingleFact)
}
