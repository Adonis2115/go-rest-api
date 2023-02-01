package main

import (
	"os"

	"github.com/Adonis2115/go-rest-api/database"
	"github.com/Adonis2115/go-rest-api/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := initApp()
	if err != nil {
		panic(err)
	}
	defer database.CloseMongoDB()
	app := generateApp()
	port := os.Getenv("PORT")
	app.Listen(":" + port)
}

func generateApp() *fiber.App {
	app := fiber.New()
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	libGroup := app.Group("/library")
	libGroup.Get("/", handlers.GetLibraries)
	libGroup.Post("/", handlers.CreateLibrary)
	return app
}

func initApp() error {
	err := loadENV()
	if err != nil {
		return err
	}
	err = database.StartMongoDB()
	if err != nil {
		return err
	}
	return nil
}

func loadENV() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
