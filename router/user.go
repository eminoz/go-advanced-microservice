package router

import (
	"github.com/eminoz/go-redis-project/api"
	"github.com/eminoz/go-redis-project/cache"
	"github.com/eminoz/go-redis-project/repository"
	"github.com/eminoz/go-redis-project/service"
	"github.com/gofiber/fiber/v2"
)

func Setup() *fiber.App {
	f := fiber.New()
	client := cache.InitRedis()
	userCache := cache.UserCache{Redis: client}
	userCollectionSetting := repository.UserCollectionSetting()
	userService := service.UserService{UserRepository: userCollectionSetting, UserRedis: &userCache}
	controller := api.UserController{UserController: &userService}
	f.Post("/createUser", controller.CreateUser)
	f.Get("/getUserByEmail/:email", controller.GetUserByEmail)
	f.Get("/getAllUser", controller.GetAllUser)
	f.Delete("/deleteUserByEmail/:email", controller.DeleteUserByEmail)
	return f
}
