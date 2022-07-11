package model

import (
	"context"

	"github.com/go-redis/redis/v8"

	"blog/config"
	"blog/util/log"
)

// 使用文档 https://redis.uptrace.dev/#executing-commands

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(
		&redis.Options{
			Addr:     config.C.Redis.Addr,
			Password: config.C.Redis.Password,
			DB:       config.C.Redis.DB,
		},
	)
	ctx := context.Background()
	err := redisClient.Ping(ctx).Err()
	if err != nil {
		log.Logger.Error(err)
	} else {
		log.Logger.Info("Redis connected successfully!")
	}
}
