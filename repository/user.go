package repository

import (
	"fmt"
	"github.com/eminoz/go-redis-project/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type IUserRepository interface {
	CreateUser(ctx *fiber.Ctx, user *model.User) (interface{}, error)
	GetUserByEmail(ctx *fiber.Ctx, email string) model.UserDal
	GetAllUser(ctx *fiber.Ctx) []model.UserDal
}

func (u *UserCollection) CreateUser(ctx *fiber.Ctx, user *model.User) (interface{}, error) {
	insertOne, err := u.Collection.InsertOne(ctx.Context(), user)
	if err != nil {
		return nil, err

	}
	return insertOne, nil
}

func (u *UserCollection) GetUserByEmail(ctx *fiber.Ctx, email string) model.UserDal {
	filter := bson.D{{"email", email}}
	var userDal model.UserDal
	u.Collection.FindOne(ctx.Context(), filter).Decode(&userDal)
	return userDal
}
func (u UserCollection) GetAllUser(ctx *fiber.Ctx) []model.UserDal {
	find, err := u.Collection.Find(ctx.Context(), bson.D{})
	if err != nil {
		fmt.Println(err)
	}
	var user []model.UserDal
	err = find.All(ctx.Context(), &user)
	if err != nil {
		fmt.Println(err)
	}
	return user
}
