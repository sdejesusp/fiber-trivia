package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sdejesusp/fiber-trivia/database"
)

func main() {
    database.ConnectDb()
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, from a Docker fiber container.")
    })

    app.Listen(":3000")
}