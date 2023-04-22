package cache

import (
	"Blog/core/config"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRedis(t *testing.T) {
	conf, err := config.LoadConfig("../...")
	require.Error(t, err)

	rdb := InitRedis(conf)

	ctx := context.Background()

	_, err = rdb.Ping(ctx).Result()
	require.NoError(t, err)

	err = rdb.Set(ctx, "key", "val", time.Minute).Err()
	require.NoError(t, err)

	val, err := rdb.Get(ctx, "key").Result()
	require.NoError(t, err)
	require.Equal(t, val, "val")

	_, err = rdb.Get(ctx, "key2").Result()
	require.Error(t, err)

}
