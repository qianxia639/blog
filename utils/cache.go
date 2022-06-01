package utils

import (
	"context"
	"time"

	"github.com/qianxia/blog/global"
)

var ctx = context.Background()

func SetCache(key string, val interface{}) error {
	return global.QX_REDIS.Set(ctx, key, val, 10*time.Minute).Err()
}

func GetCache(key string, val interface{}) error {
	return global.QX_REDIS.Get(ctx, key).Scan(val)
}

func DelCache(key string) error {
	global.QX_REDIS.Del(ctx, key)
	return nil
}

func ExisKey(key string) bool {
	return global.QX_REDIS.Exists(ctx, key).Val() > 0
}

func SetTtlCache(key string, val interface{}, ttl time.Duration) error {
	return global.QX_REDIS.SetNX(ctx, key, val, ttl).Err()
}
