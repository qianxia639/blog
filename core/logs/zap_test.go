package logs

import (
	"Blog/core/config"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestZap(t *testing.T) {

	conf, err := config.LoadConfig("../..")
	require.NoError(t, err)

	l1 := GetInstance(conf.Zap)
	l2 := GetInstance(conf.Zap)
	Logs = GetInstance(conf.Zap)
	require.Equal(t, l1, l2)
	require.Equal(t, l1, Logs)
	require.Equal(t, l2, Logs)

	t.Logf("l1: %p, l2: %p, Logs: %p", l1, l2, Logs)
}
