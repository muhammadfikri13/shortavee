package redis

import (
	"context"
	"os"

	goredis "github.com/redis/go-redis/v9"
)

var Client *goredis.Client

func Connect() {
	Client = goredis.NewClient(&goredis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
