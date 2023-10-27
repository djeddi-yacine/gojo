package worker

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const TaskDeleteSession = "task:delete_session"

type PayloadDeleteSession struct {
	Username string `json:"username"`
}

func (distributor *RedisTaskDistributor) DistributeTaskDeleteSession(
	ctx context.Context,
	payload *PayloadDeleteSession,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	task := asynq.NewTask(TaskDeleteSession, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")
	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskDeleteSession(
	ctx context.Context,
	task *asynq.Task,
) error {
	var payload PayloadDeleteSession
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	err := processor.gojo.RefreshSessions(ctx, payload.Username)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return fmt.Errorf("session doesn't exist: %w", asynq.SkipRetry)
		}
		return fmt.Errorf("failed to delete session: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).Msg("processed task")

	return nil
}
