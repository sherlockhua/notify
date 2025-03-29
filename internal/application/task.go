package application

import (
	"context"
	"notify/internal/domain/task"
)

type Appalition interface {
	CreateTask(ctx context.Context, userId int64, task *task.Task) error
}

type AppalitionImpl struct {
	ctx         context.Context
	taskService task.TaskService
}

func NewAppalition(ctx context.Context, taskService task.TaskService) Appalition {
	return &AppalitionImpl{
		ctx:         ctx,
		taskService: taskService,
	}
}

func (a *AppalitionImpl) CreateTask(ctx context.Context, userId int64, task *task.Task) error {
	return a.taskService.CreateTask(a.ctx, userId, task)
}
