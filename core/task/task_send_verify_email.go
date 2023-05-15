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
