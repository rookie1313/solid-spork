package api

import (
	"github.com/gofiber/fiber/v2"
	"solid-spork/src/dtos"
	"solid-spork/src/model"
)

func (server *Server) UserRoute(router fiber.Router) {
	router.Post("/", server.createUser)
}

func (server *Server) createUser(ctx *fiber.Ctx) error {
	req := new(dtos.CreateUserRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	e := &model.User{Email: req.Email, Password: req.Password}
	result := server.DB.Create(e)
	if result.Error != nil {
		return result.Error
	}

	return nil
}