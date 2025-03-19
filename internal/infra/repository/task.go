package repository

import (
	"context"

	"github.com/sherlockhua/koala/logs"
	"github.com/sherlockhua/notify/internal/domain/task"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB //+
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetTask(ctx context.Context, userId int64, taskId int64) (*task.Task, error) { //-
	task := &task.Task{}
	err := r.db.Where("id = ? and user_id = ?", taskId, userId).First(task).Error
	if err != nil {
		logs.Errorf(ctx, "get task failed, err:%v", err)
		return nil, err
	}
	return task, nil
}

func (r *TaskRepository) UpdateTask(ctx context.Context, userId int64, taskId int64, task *task.Task) error {
	err := r.db.Model(task).Where("id = ? and user_id = ?", taskId, userId).Updates(task).Error
	if err != nil {
		logs.Errorf(ctx, "update task failed, err:%v", err)
		return err
	}
	return nil
}

func (r *TaskRepository) CreateTask(ctx context.Context, userId int64, task *task.Task) error {
	err := r.db.Create(task).Error
	if err != nil {
		logs.Errorf(ctx, "create task failed, err:%v", err)
		return err
	}
	return nil
}

func (r *TaskRepository) DeleteTask(ctx context.Context, userId int64, taskId int64) error {
	err := r.db.Where("id = ? and user_id = ?", taskId, userId).Delete(&task.Task{}).Error
	if err != nil {
		logs.Errorf(ctx, "delete task failed, err:%v", err)
		return err
	}
	return nil
}

func (r *TaskRepository) GetTaskList(ctx context.Context, userId int64, offset, size int32) ([]*task.Task, error) {
	tasks := make([]*task.Task, 0)
	err := r.db.Where("user_id = ?", userId).Offset(int(offset)).Limit(int(size)).Find(&tasks).Error
	if err != nil {
		logs.Errorf(ctx, "get task list failed, err:%v", err)
		return nil, err
	}
	return tasks, nil
}
