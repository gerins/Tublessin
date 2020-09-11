package redis

import (
	"context"
	"strconv"
	"tublessin/services/user_service/config"

	log "github.com/sirupsen/logrus"

	"github.com/go-redis/redis/v8"
)

func NewRedisConnection() *redis.Client {
	databaseSelected, err := strconv.Atoi(config.REDIS_DATABASE_SELECT)
	if err != nil {
		log.Println("Please input valid number for REDIS_DATABASE_SELECT")
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_DATABASE_HOST + ":" + config.REDIS_DATABASE_PORT,
		Username: config.REDIS_DATABASE_USERNAME,
		Password: config.REDIS_DATABASE_PASSWORD,
		DB:       databaseSelected,
	})

	log.Println("Trying connect to Redis database...................")
	pong, err := rdb.Ping(context.Background()).Result()
	log.Println(pong, err)

	return rdb
}
