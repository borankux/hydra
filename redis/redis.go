package redis

import (
	"github.com/go-redis/redis/v8"
	"time"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		Password:    "",
		DB:          0,
		ReadTimeout: time.Second * 5,
	})
}

func GetRedis() *redis.Client {
	return rdb
}
