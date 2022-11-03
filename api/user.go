package api

import (
	"github.com/eminoz/go-advanced-microservice/service"
	"github.com/gofiber/fiber/v2"
)

type IUserController interface {
	CreateUser(ctx *fiber.Ctx) error
	GetUserByEmail(ctx *fiber.Ctx) error
	DeleteUserByEmail(ctx *fiber.Ctx) error
	UpdatedUserByEmail(ctx *fiber.Ctx) error
	SignIn(ctx *fiber.Ctx) error
	GetAllUser(ctx *fiber.Ctx) error
	GetUserAddress(ctx *fiber.Ctx) error
	CreateAddress(ctx *fiber.Ctx) error
}
type UserController struct {
	UserController service.IUserService
}

func NewUserController(u service.IUserService) IUserController {
	return &UserController{
		UserController: u,
	}
}
func (u UserController) SignIn(ctx *fiber.Ctx) error {
	in, resultError := u.UserController.SignIn(ctx)
	if resultError != nil {
		return ctx.JSON(resultError)
	}
	return ctx.JSON(in)

}
func (u UserController) CreateAddress(ctx *fiber.Ctx) error {
	address := u.UserController.CreateAddress(ctx)
	return ctx.JSON(address)
}
func (u UserController) GetUserAddress(ctx *fiber.Ctx) error {
	address := u.UserController.GetUsersAddress(ctx)
	return ctx.JSON(address)

}

func (u UserController) CreateUser(ctx *fiber.Ctx) error {
	createUser, resultError := u.UserController.CreateUser(ctx)
	if resultError != nil {
		return ctx.JSON(resultError)
	}
	return ctx.JSON(createUser)
}
func (u UserController) GetUserByEmail(ctx *fiber.Ctx) error {
	email, resultError := u.UserController.GetUserByEmail(ctx)
	if resultError != nil {
		return ctx.JSON(resultError)
	}
	return ctx.JSON(email)
}
func (u UserController) GetAllUser(ctx *fiber.Ctx) error {
	user := u.UserController.GetAllUser(ctx)
	return ctx.JSON(user)

}
func (u UserController) DeleteUserByEmail(ctx *fiber.Ctx) error {
	email, resultError := u.UserController.DeleteUserByEmail(ctx)
	if resultError != nil {
		return ctx.JSON(resultError)
	}
	return ctx.JSON(email)

}
func (u UserController) UpdatedUserByEmail(ctx *fiber.Ctx) error {
	email, resultError := u.UserController.UpdateUserByEmail(ctx)
	if resultError != nil {
		return ctx.JSON(resultError)
	}
	return ctx.JSON(email)
}
