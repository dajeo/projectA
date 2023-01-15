package db

import (
	"github.com/go-redis/redis/v9"
	"projectA/config"
)

var rdb *redis.Client

func InitRedis() {
	c := config.GetConfig()
	rdb = redis.NewClient(&redis.Options{
		Addr:     c.GetString("redis.host"),
		Password: c.GetString("redis.pass"),
		DB:       0,
	})
}

func GetRedis() *redis.Client {
	return rdb
}
