package entity

import (
	"context"
	"notify/internal/common"
	"notify/internal/domain/entity/task_strategy"
	"time"
)

type TaskTemplate struct {
	TemplateName  string               `json:"template_name"`
	TaskDesc      string               `json:"task_desc"`
	CreateTime    time.Time            `json:"create_time"`
	UpdateTime    time.Time            `json:"update_time"`
	TaskType      common.TaskType      `json:"task_type"`
	NotifyChannel common.NotifyChannel `json:"notify_channel"`
	TaskStatus    common.TaskStatus    `json:"task_status"`
	StrategyData  string               `json:"strategy_data"`
}

func (t *TaskTemplate) ConvertToTask(ctx context.Context, factory task_strategy.TaskStrategyFactory) (*Task, error) {
	task := &Task{
		TaskName:      t.TemplateName,
		TaskDesc:      t.TaskDesc,
		TaskType:      t.TaskType,
		NotifyChannel: t.NotifyChannel,
		TaskStatus:    t.TaskStatus,
		StrategyData:  t.StrategyData,
	}

	strategy, err := factory.CreateTaskStrategy(ctx, t.TaskType, t.StrategyData)
	if err != nil {
		return nil, err
	}
	task.Strategy = strategy
	return task, nil
}
