package cache

import (
	"context"
	"encoding/json"
	"github.com/eminoz/go-redis-project/model"
	"github.com/go-redis/redis/v8"
)

type IUserCache interface {
	SaveUserByEmail(user model.UserDal) error
	GetUserByEmail(email string) model.UserDal
}

type UserCache struct {
	Redis *redis.Client
}

func (c *UserCache) SaveUserByEmail(user model.UserDal) error {
	ctx := context.TODO()

	marshal, _ := json.Marshal(user)
	c.Redis.HSet(ctx, "users", user.Email, marshal)

	/*type Author struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	author := Author{Name: "emin", Age: 25}
	user, err := json.Marshal(author)
	if err != nil {
		fmt.Println(err)
	}
	client.HSet(ctx, "users", author.Name, user)
	getAll := client.HGet(ctx, "users", author.Name)
	var a Author
	json.Unmarshal([]byte(getAll.Val()), &a)
	fmt.Println(a.Age)*/

	return nil
}
func (c *UserCache) GetUserByEmail(email string) model.UserDal {
	ctx := context.TODO()
	hGet := c.Redis.HGet(ctx, "users", email)
	var user model.UserDal
	json.Unmarshal([]byte(hGet.Val()), &user)
	return user
}
