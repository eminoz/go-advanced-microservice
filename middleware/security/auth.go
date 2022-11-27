package security

import (
	"fmt"
	"github.com/eminoz/go-advanced-microservice/config"
	"github.com/eminoz/go-advanced-microservice/core/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func UserIsAuth() fiber.Handler {

	return func(ctx *fiber.Ctx) error {

		token, err := checkRole(ctx)
		if err != nil {
			ctx.JSON(*utilities.ErrorResult("Your Token has been expired"))
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "user" || claims["role"] == "admin" {

				ctx.Request().Header.Set("Role", "user")
				ctx.Next()

			}
		}
		return nil
	}
}
func AdminIsAuth() fiber.Handler {

	return func(ctx *fiber.Ctx) error {

		token, err := checkRole(ctx)
		if err != nil {
			ctx.JSON(*utilities.ErrorResult("Your Token has been expired"))
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {

				ctx.Request().Header.Set("Role", "admin")
				ctx.Next()

			}
		}
		return nil
	}
}

func checkRole(ctx *fiber.Ctx) (*jwt.Token, error) {
	header := ctx.GetReqHeaders()["Token"]
	if header == "" {
		ctx.JSON(*utilities.ErrorResult("please sign in"))
	}
	var mySigningKey = []byte(config.GetConfig().AppSecret)
	token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return mySigningKey, nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}
