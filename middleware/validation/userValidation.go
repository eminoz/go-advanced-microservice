package validation

import (
	"github.com/eminoz/go-advanced-microservice/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserValidation() fiber.Handler {
	u := new(model.User)
	validate := validator.New()
	return func(ctx *fiber.Ctx) error {
		ctx.BodyParser(u)
		err := validate.Struct(u)
		if err != nil {

			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Next()
	}

}
