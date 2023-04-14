package redis

import (
	"errors"
	"time"

	database "github.com/Benny66/ginServer/db"
	redis2 "github.com/go-redis/redis/v8"
)

func BatchPushQueue(queueName string, value interface{}) (err error) {
	_, err = database.RRedisClient.RedisCli.LPush(database.RRedisClient.Ctx, queueName, value).Result()
	return
}

func PopQueue(queueName string, timeout time.Duration) (data string, err error) {
	nameAndData, err := database.RRedisClient.RedisCli.BRPop(database.RRedisClient.Ctx, timeout, []string{queueName}...).Result()
	if err == redis2.Nil {
		return "", errors.New("查询不到数据")
	}
	if err != nil {
		// 查问出错
		return "", err
	}
	if len(nameAndData) > 1 {
		data = nameAndData[1]
	}
	return
}
