package redis

import (
	"errors"

	redis "github.com/go-redis/redis/v8"
)

/*
	向 redis hash 中存值
	参数：
	key：存入hash的key值
	mapData：对应key的map值
	mapData 格式： map[string]interface{}{"key1": "value1", "key2": "value2"}
	return：bool（是否添加成功），error（错误信息）
*/
func (r *RRedis) HashSet(key string, mapData map[string]interface{}) (bool, error) {
	// 参数非空验证
	if key == "" || mapData == nil {
		return false, errors.New("参数为空")
	}
	if r.RedisCli == nil {
		return false, errors.New("客户端断开连接")
	} else {
		if err := r.RedisCli.HSet(r.Ctx, key, mapData).Err(); err != nil {
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
func (r *RRedis) HashGet(key string, field string) (string, error) {
	// 参数非空判断
	if key == "" || field == "" {
		return "", errors.New("参数为空")
	}
	value, err := r.RedisCli.HGet(r.Ctx, key, field).Result()
	if err == redis.Nil {
		return "", errors.New("key 不存在")
	} else if err != nil {
		return "", errors.New("获取失败")
	}
	return value, nil
}

/*
	根据key，field 获取值
	参数：
	key：存入hash的key值
	fields：可变长参数，0到n个field
	return：map[string]interface{} 返回一个map
*/
func (r *RRedis) BatchHashGet(key string, fields ...string) ([]interface{}, error) {
	if key == "" {
		return nil, errors.New("参数为空")
	}
	resultArray, err := r.RedisCli.HMGet(r.Ctx, key, fields...).Result()
	if err != nil {
		return nil, errors.New("error occur when get data from redis : " + err.Error())
	}
	return resultArray, nil
}

/*
	判断 hash key，field是否存在
	参数：
	key：存入hash的key值
	field：字段名
	返回值：
	bool：字段是否存在
	error：错误信息
*/
func (r *RRedis) HashKeyExist(key, field string) (bool, error) {
	if key == "" || field == "" {
		return false, errors.New("参数为空")
	}
	b, err := r.RedisCli.HExists(r.Ctx, key, field).Result()
	if err != nil {
		return false, errors.New("异常：" + err.Error())
	}
	return b, nil
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
func (r *RRedis) HashDel(key string, fields ...string) (bool, error) {
	if key == "" {
		return false, errors.New("参数为空")
	}
	_, err := r.RedisCli.HDel(r.Ctx, key, fields...).Result()
	if err != nil {
		return false, errors.New("异常：" + err.Error())
	}
	return true, nil
}
