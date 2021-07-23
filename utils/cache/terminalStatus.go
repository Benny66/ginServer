package cache

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-06-18 10:32:00
 * @LastEditors: shahao
 * @LastEditTime: 2021-06-18 11:17:15
 */

import (
	"time"

	goChache "github.com/patrickmn/go-cache"
)

var TerminalStatusCache *terminalStatusCache

func init() {
	// 创建一个默认过期时间为3s的缓存适配器
	// 每1清除一次过期的项目
	TerminalStatusCache = NewTerminalStatusCache()
}

func NewTerminalStatusCache() *terminalStatusCache {
	return &terminalStatusCache{
		Cache: goChache.New(10*time.Second, 5*time.Second),
	}
}

// var terminalStatusCache *goChache.Cache

type terminalStatusCache struct {
	Cache *goChache.Cache
}

func (cache *terminalStatusCache) SetCahce(k string, x interface{}, d time.Duration) {
	cache.Cache.Set(k, x, d)
}

func (cache *terminalStatusCache) GetCache(k string) (interface{}, bool) {
	return cache.Cache.Get(k)
}

//设置cache 无时间参数
func (cache *terminalStatusCache) SetDefaultCahce(k string, x interface{}) {
	cache.Cache.SetDefault(k, x)
}

//删除 cache
func (cache *terminalStatusCache) DeleteCache(k string) {
	cache.Cache.Delete(k)
}

// Add() 加入缓存
func (cache *terminalStatusCache) AddCache(k string, x interface{}, d time.Duration) {
	cache.Cache.Add(k, x, d)
}

// IncrementInt() 对已存在的key 值自增n
func (cache *terminalStatusCache) IncrementIntCahce(k string, n int) (num int, err error) {
	return cache.Cache.IncrementInt(k, n)
}
