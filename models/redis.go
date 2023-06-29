package models

import (
	"context"
	"myGo/common/logger"
	//"github.com/go-redis/redis"
	"github.com/go-redis/redis/v8"
	"myGo/config"
	"time"
)

var RedisClient *redis.Client
var ctx = context.Background()

// 初始化
func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         config.RedConf.Addr + ":" + config.RedConf.Port, // 连接地址
		Password:     config.RedConf.Password,                         // 密码
		DB:           config.RedConf.DB,                               // 数据库编号
		PoolSize:     config.RedConf.PoolSize,
		MaxRetries:   config.RedConf.MaxRetries,
		DialTimeout:  time.Duration(config.RedConf.DialTimeout) * time.Millisecond, // 链接超时
		ReadTimeout:  time.Duration(config.RedConf.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(config.RedConf.WriteTimeout) * time.Millisecond,
	})
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		logger.Error("Redis初始化失败:", err.Error())
		panic("Redis连接失败:" + err.Error())
	}

	logger.Info("Redis初始化完成......")
}
