package api

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"solid-spork/src/dtos"
	"solid-spork/src/model"
)

func (server *Server) UserRoute(router fiber.Router) {
	router.Post("/", server.createUser)
	router.Get("/usersList", server.getAllUsers)
	router.Get("/", server.getUserByEmail)
}

func (server *Server) createUser(ctx *fiber.Ctx) error {
	req := new(dtos.CreateUserRequest)
	if err := ctx.BodyParser(req); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}
	if err := validate.Struct(req); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}
	e := &model.User{Email: req.Email, Password: req.Password, Name: req.Name}
	result := server.DB.Create(e)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return &fiber.Error{
				Code:    fiber.StatusConflict,
				Message: result.Error.Error(),
			}
		}
		return result.Error
	}

	response := &dtos.UserResponse{
		ID:    e.Model.ID,
		Name:  e.Name,
		Email: e.Email,
	}
	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (server *Server) getAllUsers(ctx *fiber.Ctx) error {
	var response []dtos.UserResponse
	result := server.DB.Model(&model.User{}).Find(&response)
	if result.Error != nil {
		return result.Error
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

type findUserRequest struct {
	Email string `query:"email"`
}

func (server *Server) getUserByEmail(ctx *fiber.Ctx) error {
	u := new(findUserRequest)
	if err := ctx.QueryParser(u); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}
	var response dtos.UserResponse
	result := server.DB.Model(&model.User{}).Where("email = ?", u.Email).First(&response)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &fiber.Error{
				Code:    fiber.StatusOK,
				Message: result.Error.Error(),
			}
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
