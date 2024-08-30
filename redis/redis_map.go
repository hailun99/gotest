package redis

import (
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "addr:1926",
		Password: "",
		DB:       0,
	})
}
