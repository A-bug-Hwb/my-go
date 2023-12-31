package redisUtil

import (
	"context"
	"myGo/models"
	"time"
)

// RedisUtil 公用
var RedisUtil = &models.RedisClient
var ctx = context.Background()

// Set 插入
func Set(key string, value interface{}) bool {
	_, err := (*RedisUtil).GetSet(ctx, key, value).Result()
	if err != nil {
		return false
	}
	return true
}

// SetTime 插入（带过期时间）
func SetTime(key string, value interface{}, expiration time.Duration) {
	(*RedisUtil).Set(ctx, key, value, expiration)
}

// Get 获取
func Get(key string) interface{} {
	res, err := (*RedisUtil).Get(ctx, key).Result()
	if err != nil {
		return nil
	}
	return res
}

// Del 删除
func Del(key ...string) bool {
	_, err := (*RedisUtil).Del(ctx, key...).Result()
	if err != nil {
		return false
	}
	return true
}
