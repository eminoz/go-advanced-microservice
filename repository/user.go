package repository

import (
	"fmt"
	"github.com/eminoz/go-advanced-microservice/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type IUserRepository interface {
	CreateUser(ctx *fiber.Ctx, user *model.User) (model.UserDal, error)
	GetUserByEmail(ctx *fiber.Ctx, email string) model.UserDal
	GetAllUser(ctx *fiber.Ctx) []model.UserDal
	DeleteUserByEmail(ctx *fiber.Ctx, email string) (int64, error)
	UpdateUserByEmail(ctx *fiber.Ctx, email string, user model.UserDal) (bool, string)
	GetUserByEmailForAuth(ctx *fiber.Ctx, email string) model.User
}

func (u UserCollection) CreateUser(ctx *fiber.Ctx, user *model.User) (model.UserDal, error) {
	insertOne, err := u.Collection.InsertOne(ctx.Context(), user)
	var userDal model.UserDal
	if err != nil {
		return userDal, err
	}
	filter := bson.D{{"_id", insertOne.InsertedID}}

	u.Collection.FindOne(ctx.Context(), filter).Decode(&userDal)
	return userDal, nil
}

func (u UserCollection) GetUserByEmailForAuth(ctx *fiber.Ctx, email string) model.User {
	filter := bson.D{{"email", email}}
	var User model.User
	u.Collection.FindOne(ctx.Context(), filter).Decode(&User)
	return User
}
func (u UserCollection) GetUserByEmail(ctx *fiber.Ctx, email string) model.UserDal {
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
func (u UserCollection) DeleteUserByEmail(ctx *fiber.Ctx, email string) (int64, error) {
	d := bson.D{{"email", email}}
	deleteOne, err := u.Collection.DeleteOne(ctx.Context(), d)
	if err != nil {
		return 0, err
	}
	return deleteOne.DeletedCount, nil
}
func (u UserCollection) UpdateUserByEmail(ctx *fiber.Ctx, email string, user model.UserDal) (bool, string) {
	filter := bson.D{{"email", email}}
	var usr model.User
	err := u.Collection.FindOne(ctx.Context(), filter).Decode(&usr)
	if err != nil {
		return false, "user did not found"
	}
	usr.Email = user.Email
	usr.Name = user.Name
	update := bson.D{{"$set", usr}}
	updateOne, _ := u.Collection.UpdateOne(ctx.Context(), filter, update)
	if updateOne.ModifiedCount > 0 {
		return true, "user updated"
	}
	return false, "user did not updated"

}
