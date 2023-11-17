package api

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	*fiber.App
	*gorm.DB
}

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

var validate = validator.New(validator.WithRequiredStructEnabled())

func CreateServer(db *gorm.DB) *Server {
	app := fiber.New(fiber.Config{
		Immutable: true,
		// glob error handle
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			return ctx.Status(code).JSON(GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		}})
	server := Server{app, db}

	apiV1 := app.Group("/api/v1")
	apiV1.Route("/user", server.UserRoute)
	return &server
}
