package router

import (
	"github.com/eminoz/go-advanced-microservice/api"
	"github.com/eminoz/go-advanced-microservice/cache"
	"github.com/eminoz/go-advanced-microservice/repository"
	"github.com/eminoz/go-advanced-microservice/security/encryption"
	"github.com/eminoz/go-advanced-microservice/security/jwt"
	"github.com/eminoz/go-advanced-microservice/service"
)

type base struct{}

var client = cache.InitRedis()
var auth = jwt.Auth{}
var userEncryption = encryption.UserEncryption{}
var userCache = cache.UserCache{Redis: client}

func (b base) UserDI() api.IUserController {
	//DI for user service
	userCollection := repository.UserCollectionSetting()
	newUserService := service.NewUserService(userCollection, userCache, auth, userEncryption)
	userController := api.NewUserController(newUserService)
	return userController
}
func (b base) OrderDI() api.IOrderController {
	//DI for order service
	userCollection := repository.UserCollectionSetting()
	orderService := service.NewOrderService(userCollection)
	orderController := api.NewOrderController(orderService)
	return orderController
}
