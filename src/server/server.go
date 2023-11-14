package server

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateApp(db *gorm.DB) *fiber.App {
	app := fiber.New(fiber.Config{Immutable: true})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	return app
}
