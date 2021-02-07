package db

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var Redis *redis.Client

func InitRedis(addr string) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
