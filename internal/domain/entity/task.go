package entity

import (
	"context"
	"notify/internal/common"
	"notify/internal/domain/entity/task_strategy"
	"time"
)

type Task struct {
	TaskID        int64                      `json:"task_id"`
	UserID        int64                      `json:"user_id"`
	TaskName      string                     `json:"task_name"`
	TaskDesc      string                     `json:"task_desc"`
	CreateTime    time.Time                  `json:"create_time"`
	UpdateTime    time.Time                  `json:"update_time"`
	TaskType      common.TaskType            `json:"task_type"`
	NotifyChannel common.NotifyChannel       `json:"notify_channel"`
	TaskStatus    common.TaskStatus          `json:"task_status"`
	StrategyData  string                     `json:"strategy_data"`
	Strategy      task_strategy.TaskStrategy `json:"-"`
}

func (t *Task) IsTimeToNotify(ctx context.Context, factory task_strategy.TaskStrategyFactory) (common.NotifyTimeResult, error) {
	return t.Strategy.IsTimeToNotify(ctx), nil
}
