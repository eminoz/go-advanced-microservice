package api

import (
	"github.com/eminoz/go-redis-project/service"
	"github.com/gofiber/fiber/v2"
)

type IUserController interface {
	CreateUser(ctx *fiber.Ctx) error
	GetUserByEmail(ctx *fiber.Ctx) error
}
type UserController struct {
	UserController service.IUserService
}

func (u *UserController) CreateUser(ctx *fiber.Ctx) error {
	createUser, err := u.UserController.CreateUser(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(createUser)
}
func (u UserController) GetUserByEmail(ctx *fiber.Ctx) error {
	email := u.UserController.GetUserByEmail(ctx)
	return ctx.JSON(email)
}
func (u UserController) GetAllUser(ctx *fiber.Ctx) error {
	user := u.UserController.GetAllUser(ctx)
	return ctx.JSON(user)

}