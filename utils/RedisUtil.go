package utils

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/qianxia/blog/config"
)

var Pool *redis.Pool

func InitRedis(y *config.Config) *redis.Pool {
	pool := &redis.Pool{ // 实例化一个链接池
		MaxIdle:     8,   // 最初的链接数
		MaxActive:   0,   // 最大连接数
		IdleTimeout: 300, // 链接关闭时间
		Dial: func() (redis.Conn, error) { // 要链接的redis数据库
			return redis.Dial("tcp", fmt.Sprintf("%s:%d", y.Redis.Host, y.Redis.Port))
		},
	}
	Pool = pool
	return pool
}

func GetRedisConn() redis.Conn {
	return Pool.Get()
}
