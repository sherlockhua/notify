package repository

import (
	"context"
	"notify/internal/common"
	"notify/internal/domain/entity"
)

type TaskRepository interface {
	GetTask(ctx context.Context, userId int64, taskId int64) (*entity.Task, error)
	UpdateTask(ctx context.Context, userId int64, task *entity.Task) error
	CreateTask(ctx context.Context, userId int64, task *entity.Task) error
	DeleteTask(ctx context.Context, userId int64, taskId int64) error
	GetTaskList(ctx context.Context, userId int64, offset, size int32,
		conditions map[string]interface{}, statusConditions []common.TaskStatus) ([]*entity.Task, error)

	GetAllTaskList(ctx context.Context, offset, size int32,
		conditions map[string]interface{}, statusConditions []common.TaskStatus) ([]*entity.Task, error)

	CreateTaskTemplate(ctx context.Context, template *entity.TaskTemplate) error
}
