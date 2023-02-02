package redis

import (
	"errors"
	"time"

	redis2 "github.com/go-redis/redis/v8"
)

func (r *RRedis) BatchPushQueue(queueName string, value interface{}) (err error) {
	_, err = r.RedisCli.LPush(r.Ctx, queueName, value).Result()
	return
}

func (r *RRedis) PopQueue(queueName string, timeout time.Duration) (data string, err error) {
	nameAndData, err := r.RedisCli.BRPop(r.Ctx, timeout, []string{queueName}...).Result()
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
