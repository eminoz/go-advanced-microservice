package router

import (
	"github.com/eminoz/go-advanced-microservice/middleware/security"
	"github.com/eminoz/go-advanced-microservice/middleware/validation"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Setup() *fiber.App {
	f := fiber.New()
	f.Use(cors.New())
	userDI := base{}
	var u = userDI.UserDI()
	var o = userDI.OrderDI()
	f.Post("/createUser", validation.UserValidation(), u.CreateUser)
	f.Post("/signin", u.SignIn)
	f.Put("/updateUser/:email", security.IsAuth(), u.UpdatedUserByEmail)
	f.Get("/getUserByEmail/:email", u.GetUserByEmail)
	f.Get("/getAllUser", u.GetAllUser)
	f.Delete("/deleteUserByEmail/:email", u.DeleteUserByEmail)
	//DI for order service
	f.Post("/createOrder/:id", o.CreateOrder)
	return f
}
