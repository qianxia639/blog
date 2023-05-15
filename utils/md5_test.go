package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var path = "/opt/blog/sqlc.yaml"

func TestShardMd5(t *testing.T) {

	hash, err := ShardMd5(path)
	require.NoError(t, err)
	t.Logf("hash: %s\n", hash)
}

func BenchmarkShardMd5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ShardMd5(path)
		require.NoError(b, err)
	}
}

func BenchmarkShardMd5Single(b *testing.B) {
	_, err := ShardMd5(path)
	require.NoError(b, err)
}
