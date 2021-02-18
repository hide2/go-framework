package db

import (
	"context"

	. "server/pkg/config"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var Redis *redis.Client

func init() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     GlobalConfig.Redis,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
