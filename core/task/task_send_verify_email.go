package task

import (
	"Blog/core/logs"
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

const (
	TaskSendVerifyEmail = "task:send_verify_email"
)

type SendVerifyEmailPayload struct {
	Email string
}

func (distributor *RedisTaskDistributor) DistributeTaskSendVerifyEmail(ctx context.Context, payload *SendVerifyEmailPayload, opts ...asynq.Option) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	tak := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)
	taskInfo, err := distributor.client.EnqueueContext(ctx, tak)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}

	logs.Logs.Info("enqueued task", zap.String("type", tak.Type()), zap.ByteString("payload", taskInfo.Payload), zap.String("queue", taskInfo.Queue), zap.Int("max_retry", taskInfo.MaxRetry))

	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, tak *asynq.Task) error {
	var payload SendVerifyEmailPayload
	err := json.Unmarshal(tak.Payload(), &payload)
	if err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	// TODO: send email to user
	logs.Logs.Info("processed task", zap.String("type", tak.Type()), zap.ByteString("payload", tak.Payload()), zap.String("email", payload.Email))

	return nil
}
