package service

import (
	"context"
	"notify/internal/common"
	"notify/internal/domain/entity"
	"notify/internal/domain/repository"

	"github.com/sherlockhua/koala/logs"
)

type TaskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) GetTask(ctx context.Context, userId int64, taskId int64) (*entity.Task, error) {
	logs.Debugf(ctx, "getting task, userId:%d, taskId:%d", userId, taskId)
	return s.taskRepo.GetTask(ctx, userId, taskId)
}

func (s *TaskService) UpdateTask(ctx context.Context, userId int64, task *entity.Task) error {
	logs.Debugf(ctx, "updating task, userId:%d, task:%+v", userId, task)
	return s.taskRepo.UpdateTask(ctx, userId, task)
}

func (s *TaskService) CreateTask(ctx context.Context, userId int64, task *entity.Task) error {
	logs.Debugf(ctx, "creating task, userId:%d, task:%+v", userId, task)
	task.TaskStatus = common.TaskStatusNotStart
	return s.taskRepo.CreateTask(ctx, userId, task)
}

func (s *TaskService) DeleteTask(ctx context.Context, userId int64, taskId int64) error {
	logs.Debugf(ctx, "deleting task, userId:%d, taskId:%d", userId, taskId)
	return s.taskRepo.DeleteTask(ctx, userId, taskId)
}

func (s *TaskService) GetTaskList(ctx context.Context, userId int64, offset, size int32) ([]*entity.Task, error) {
	logs.Debugf(ctx, "getting task list, userId:%d, offset:%d, size:%d", userId, offset, size)
	return s.taskRepo.GetTaskList(ctx, userId, offset, size)
}
