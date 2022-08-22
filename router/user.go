package router

import (
	"github.com/eminoz/go-advanced-microservice/api"
	"github.com/eminoz/go-advanced-microservice/cache"
	"github.com/eminoz/go-advanced-microservice/middleware/security"
	"github.com/eminoz/go-advanced-microservice/middleware/validation"
	"github.com/eminoz/go-advanced-microservice/repository"
	"github.com/eminoz/go-advanced-microservice/security/encryption"
	"github.com/eminoz/go-advanced-microservice/security/jwt"
	"github.com/eminoz/go-advanced-microservice/service"
	"github.com/gofiber/fiber/v2"
)

func Setup() *fiber.App {
	f := fiber.New()
	client := cache.InitRedis()
	auth := jwt.Auth{}
	userEncryption := encryption.UserEncryption{}
	userCache := cache.UserCache{Redis: client}

	//DI for user service
	userCollectionSetting := repository.UserCollectionSetting()
	newUserService := service.NewUserService(userCollectionSetting, userCache, auth, userEncryption)
	userController := api.NewUserController(newUserService)

	f.Post("/createUser", validation.UserValidation(), userController.CreateUser)
	f.Post("/signin", userController.SignIn)
	f.Put("/updateUser/:email", security.IsAuth(), userController.UpdatedUserByEmail)
	f.Get("/getUserByEmail/:email", userController.GetUserByEmail)
	f.Get("/getAllUser", userController.GetAllUser)
	f.Delete("/deleteUserByEmail/:email", userController.DeleteUserByEmail)

	//DI for order service
	orderService := service.NewOrderService(userCollectionSetting)
	orderController := api.NewOrderController(orderService)

	f.Post("/createOrder/:id", orderController.CreateOrder)
	return f
}
