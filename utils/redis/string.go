package redis

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2022-01-14 12:05:46
 * @LastEditors: shahao
 * @LastEditTime: 2022-01-14 13:30:17
 */

import (
	"errors"
	"time"

	database "github.com/Benny66/ginServer/db"
	"github.com/go-redis/redis/v8"
)

func Set(key string, data interface{}, expiration time.Duration) (bool, error) {
	// 参数非空验证
	if key == "" || data == "" {
		return false, errors.New("参数为空")
	}
	if database.RRedisClient.RedisCli == nil {
		return false, errors.New("客户端断开连接")
	} else {
		if err := database.RRedisClient.RedisCli.Set(database.RRedisClient.Ctx, key, data, expiration).Err(); err != nil {
			return false, errors.New("添加失败")
		} else {
			return true, nil
		}
	}
}

/*
根据key，field 获取值
参数：
key：存入hash的key值
field：字段名
return：string（返回字段的值），error（错误信息）
*/
func Get(key string) (string, error) {
	// 参数非空判断
	if key == "" {
		return "", errors.New("参数为空")
	}
	value, err := database.RRedisClient.RedisCli.Get(database.RRedisClient.Ctx, key).Result()
	if err == redis.Nil {
		return "", errors.New("key 不存在")
	} else if err != nil {
		return "", errors.New("获取失败")
	}
	return value, nil
}

/*
删除hash field
参数：
key：存入hash的key值
fields：字段名 数组
返回值：
bool：字段是否存在
error：错误信息
*/
func Del(key ...string) (bool, error) {
	if len(key) == 0 {
		return false, errors.New("参数为空")
	}
	_, err := database.RRedisClient.RedisCli.Del(database.RRedisClient.Ctx, key...).Result()
	if err != nil {
		return false, errors.New("异常：" + err.Error())
	}
	return true, nil
}
