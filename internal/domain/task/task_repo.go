package task

import (
	"context"
)

type TaskRepo interface {
	GetTask(ctx context.Context, userId int64, taskId int64) (*Task, error)
	UpdateTask(ctx context.Context, userId int64, task *Task) error
	CreateTask(ctx context.Context, userId int64, task *Task) error
	DeleteTask(ctx context.Context, userId int64, taskId int64) error
	GetTaskList(ctx context.Context, userId int64, offset, size int32) ([]*Task, error)
}
