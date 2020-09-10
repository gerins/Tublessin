package config

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/go-redis/redis/v8"
)

func NewRedisConnection() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:        "127.0.0.1:6379",
		Password:    "",
		DB:          0, // Default DB
		IdleTimeout: 10,
		MaxConnAge:  5,
	})

	pong, err := rdb.Ping(context.Background()).Result()
	log.Println(pong, err)

	return rdb
}
