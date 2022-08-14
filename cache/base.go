package cache

import "github.com/go-redis/redis/v8"

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
