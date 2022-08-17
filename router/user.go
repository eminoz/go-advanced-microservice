package router

import (
	"github.com/eminoz/go-advanced-microservice/api"
	"github.com/eminoz/go-advanced-microservice/cache"
	"github.com/eminoz/go-advanced-microservice/repository"
	"github.com/eminoz/go-advanced-microservice/service"
	"github.com/gofiber/fiber/v2"
)

func Setup() *fiber.App {
	f := fiber.New()
	client := cache.InitRedis()
	userCache := cache.UserCache{Redis: client}
	userCollectionSetting := repository.UserCollectionSetting()
	userService := service.UserService{UserRepository: userCollectionSetting,
		UserRedis: &userCache}
	controller := api.UserController{UserController: &userService}
	f.Post("/createUser", controller.CreateUser)
	f.Put("/updateUser/:email", controller.UpdatedUserByEmail)
	f.Get("/getUserByEmail/:email", controller.GetUserByEmail)
	f.Get("/getAllUser", controller.GetAllUser)
	f.Delete("/deleteUserByEmail/:email", controller.DeleteUserByEmail)
	return f
}
