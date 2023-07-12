package task

import (
	"Blog/core/logs"
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

type TaskDistributor interface {
	DistributeTaskSendVerifyEmail(ctx context.Context, payload interface{}, opts ...asynq.Option) error
}

type RedisTaskDistributor struct {
	client *asynq.Client
}

func NewRedisTaskDistributor(redisOpt asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(redisOpt)
	return &RedisTaskDistributor{
		client: client,
	}
}

func (distributor *RedisTaskDistributor) DistributeTaskSendVerifyEmail(ctx context.Context, payload interface{}, opts ...asynq.Option) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	tak := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)
	taskInfo, err := distributor.client.EnqueueContext(ctx, tak)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}

	logs.Logger.Info("enqueued task", zap.String("type", tak.Type()), zap.ByteString("payload", taskInfo.Payload), zap.String("queue", taskInfo.Queue), zap.Int("max_retry", taskInfo.MaxRetry))

	return nil
}
