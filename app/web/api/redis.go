package api

import (
	"fmt"
	"ginServer/utils/format"
	"ginServer/utils/log"
	mRedis "ginServer/utils/redis"

	"github.com/gin-gonic/gin"
)


var RedisApi *redisApi

func init() {
	RedisApi = NewRedisApi()
}

func NewRedisApi() *redisApi {
	return &redisApi{}
}

type redisApi struct {
}

func (api *redisApi) Test(context *gin.Context) {
	key := "redis_test"
	result, err := mRedis.RRedisClient.Set(key, 1000, 0)
	if err != nil {
		format.NewResponseJson(context).Error(51001, err.Error())
		return
	}
	if result {
		log.SystemLog("redis set success")
	}
	value, err := mRedis.RRedisClient.Get(key)
	if err != nil {
		format.NewResponseJson(context).Error(51001, err.Error())
		return
	}
	fmt.Println(value)

}
