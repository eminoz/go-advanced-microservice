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
	var p = userDI.ProductDI()
	f.Post("/createUser", validation.UserValidation(), u.CreateUser)
	f.Post("/signin", u.SignIn)
	f.Put("/updateUser/:email", security.UserIsAuth(), u.UpdatedUserByEmail)
	f.Get("/getUserByEmail/:email", security.UserIsAuth(), u.GetUserByEmail)
	f.Get("/getAllUser", u.GetAllUser)
	f.Post("/createAddress/:email", u.CreateAddress)
	f.Get("/getUserAddress/:email", u.GetUserAddress)
	f.Delete("/deleteUserByEmail/:email", u.DeleteUserByEmail)
	//DI for order service
	f.Post("/createOrder/:id", o.CreateOrder)
	f.Get("/getUserOrders/:id", o.GetOrders)

	group := f.Group("/product")

	group.Post("/create", security.AdminIsAuth(), p.CreateProduct)
	group.Post("/update/:productname", security.AdminIsAuth(), p.UpdateProductBProductName)
	group.Delete("/delete/:productname", security.AdminIsAuth(), p.DeleteProduct)
	group.Get("/getAll", p.GetAllProduct)
	return f
}
