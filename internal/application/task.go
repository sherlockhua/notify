package application

import (
	"context"
	"notify/internal/domain/entity"
	"notify/internal/domain/service"
)

type Appalition interface {
	CreateTask(ctx context.Context, userId int64, task *entity.Task) error
}

type AppalitionImpl struct {
	ctx         context.Context
	taskService service.TaskService
}

func NewAppalition(ctx context.Context, taskService service.TaskService) Appalition {
	return &AppalitionImpl{
		ctx:         ctx,
		taskService: taskService,
	}
}

func (a *AppalitionImpl) CreateTask(ctx context.Context, userId int64, task *entity.Task) error {
	return a.taskService.CreateTask(a.ctx, userId, task)
}
