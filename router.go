package main

import (
	"github.com/Adonis2115/go-rest-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func generateApp() *fiber.App {
	app := fiber.New()
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	libGroup := app.Group("/library")
	libGroup.Get("/", handlers.TestHandler)
	return app
}
