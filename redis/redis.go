package redis

import (
	"context"
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


func GetTotalKeySpace() int {
	return int(rdb.DBSize(context.Background()).Val())
}

func Scan(key string) *redis.ScanCmd{
	return rdb.Scan(context.Background(), 0, key, 100)
}