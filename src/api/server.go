package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	*fiber.App
	*gorm.DB
}

func CreateServer(db *gorm.DB) *Server {
	app := fiber.New(fiber.Config{Immutable: true})
	server := Server{app, db}

	apiV1 := app.Group("/api/v1")
	apiV1.Route("/user", server.UserRoute)
	return &server
}
