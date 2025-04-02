package repository

import (
	"context"
	"notify/internal/domain/entity"
)

type TaskRepository interface {
	GetTask(ctx context.Context, userId int64, taskId int64) (*entity.Task, error)
	UpdateTask(ctx context.Context, userId int64, task *entity.Task) error
	CreateTask(ctx context.Context, userId int64, task *entity.Task) error
	DeleteTask(ctx context.Context, userId int64, taskId int64) error
	GetTaskList(ctx context.Context, userId int64, offset, size int32) ([]*entity.Task, error)
}
