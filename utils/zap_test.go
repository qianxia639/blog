package utils

import (
	"testing"

	"go.uber.org/zap"
)

func TestZap(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("test log", zap.String("Func", "TestZap"))
}
