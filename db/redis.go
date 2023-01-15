package db

import (
	"context"
	"github.com/go-redis/redis/v9"
	"log"
	"projectA/config"
)

var ctx context.Context
var rdb *redis.Client

func InitRedis() {
	c := config.GetConfig()
	ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     c.GetString("redis.host"),
		Password: c.GetString("redis.pass"),
		DB:       0,
	})

	err := rdb.Ping(ctx).Err()
	if err != nil {
		log.Fatal("error on initializing redis: ", err)
	}
}

func GetRedis() *redis.Client {
	return rdb
}

func GetContext() context.Context {
	return ctx
}
