package api

import (
	"fmt"

	"github.com/Benny66/ginServer/log"
	"github.com/Benny66/ginServer/utils/format"
	myRedis "github.com/Benny66/ginServer/utils/redis"

	"github.com/gin-gonic/gin"
)

func (api *redisApi) Test(context *gin.Context) {
	key := "redis_test"
	result, err := myRedis.Set(key, 1000, 0)
	if err != nil {
		format.NewResponseJson(context).Error(51001, err.Error())
		return
	}
	if result {
		log.SystemLog("redis set success")
	}
	value, err := myRedis.Get(key)
	if err != nil {
		format.NewResponseJson(context).Error(51001, err.Error())
		return
	}
	fmt.Println(value)

}
