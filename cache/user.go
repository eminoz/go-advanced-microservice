package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eminoz/go-advanced-microservice/model"
	"github.com/go-redis/redis/v8"
)

//go:generate mockgen -destination=../mocks/cache/mockUsercache.go -package=cache  github.com/eminoz/go-advanced-microservice/cache IUserCache
type IUserCache interface {
	SaveUserByEmail(user model.UserDal) error
	GetUserByEmail(email string) model.UserDal
	GetAllUser() []model.UserDal
	DeleteUserByEmail(email string)
}

type UserCache struct {
	Redis *redis.Client
}

func (c UserCache) SaveUserByEmail(user model.UserDal) error {
	ctx := context.TODO()

	marshal, _ := json.Marshal(user)
	fmt.Println(user)
	c.Redis.HSet(ctx, "users", user.Email, marshal)
	return nil
}
func (c UserCache) GetUserByEmail(email string) model.UserDal {
	ctx := context.TODO()
	hGet := c.Redis.HGet(ctx, "users", email)
	var user model.UserDal
	json.Unmarshal([]byte(hGet.Val()), &user)
	return user
}
func (c UserCache) GetAllUser() []model.UserDal {
	ctx := context.TODO()
	getAll := c.Redis.HGetAll(ctx, "users")
	var user []model.UserDal
	var u model.UserDal
	for _, j := range getAll.Val() {
		json.Unmarshal([]byte(j), &u)
		user = append(user, u)
	}
	return user
}

func (c UserCache) DeleteUserByEmail(email string) {
	ctx := context.TODO()
	del := c.Redis.HDel(ctx, "users", email)
	fmt.Println(del.Val())
}
