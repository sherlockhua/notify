package entity

import (
	"context"
	"notify/internal/common"
	"notify/internal/domain/entity/task_strategy"
	"time"
)

type TaskTemplate struct {
	TemplateID    int64                 `json:"template_id"`
	TemplateName  string                `json:"template_name"`
	TaskDesc      string                `json:"template_desc"`
	TaskType      common.TaskType       `json:"task_type"`
	NotifyChannel common.NotifyChannel  `json:"notify_channel"`
	Status        common.TemplateStatus `json:"template_status"`
	StrategyData  string                `json:"strategy_data"`
	CreateTime    time.Time             `json:"create_time"`
	UpdateTime    time.Time             `json:"update_time"`
}

func (t *TaskTemplate) ConvertToTask(ctx context.Context, factory task_strategy.TaskStrategyFactory) (*Task, error) {

	strategy, err := factory.CreateTaskStrategy(ctx, t.TaskType, t.StrategyData)
	if err != nil {
		return nil, err
	}
	task := &Task{
		TaskName:      t.TemplateName,
		TaskDesc:      t.TaskDesc,
		TaskType:      t.TaskType,
		NotifyChannel: t.NotifyChannel,
		TaskStatus:    common.TaskStatusRunning,
		StrategyData:  t.StrategyData,
		Strategy:      strategy,
	}
	return task, nil
}
