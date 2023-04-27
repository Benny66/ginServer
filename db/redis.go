package database

import (
	"context"
	"time"

	"github.com/Benny66/ginServer/config"

	log2 "github.com/Benny66/ginServer/log"

	redis "github.com/go-redis/redis/v8"
)

type RRedis struct {
	RedisCli *redis.Client
	Ctx      context.Context
}

var RRedisClient = RRedis{
	Ctx: context.Background(),
}

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	RRedisClient.RedisCli = redis.NewClient(&redis.Options{
		Addr:     config.Config.RedisAddr,
		Password: config.Config.RedisPassword,
	})
	_, err := RRedisClient.RedisCli.Ping(ctx).Result()
	if err != nil {
		panic("连接redis失败" + err.Error())
	}
	log2.SystemLog("连接redis成功")
}
