package db

import (
	"github.com/go-redis/redis/v9"
	"os"
)

var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})
}

func GetRedis() *redis.Client {
	return rdb
}
