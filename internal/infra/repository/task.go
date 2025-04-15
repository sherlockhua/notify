package repository

import (
	"context"
	"notify/internal/common"
	"notify/internal/domain/entity"
	"notify/internal/domain/repository"
	"time"

	"github.com/sherlockhua/koala/logs"
	"gorm.io/gorm"
)

type TaskRepositoryImpl struct {
	db *gorm.DB //+
}
type TaskModel struct {
	ID               int64             `gorm:"column:id" json:"id"`
	TaskID           int64             `gorm:"column:task_id" json:"task_id"`
	TaskName         string            `gorm:"column:task_name" json:"task_name"`
	TaskDesc         string            `gorm:"column:task_desc" json:"task_desc"`
	CreateTime       time.Time         `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"`
	TaskType         int               `gorm:"column:task_type" json:"task_type"`
	NotifyType       int               `gorm:"column:notify_type" json:"notify_type"`
	NotifyTime       time.Time         `gorm:"column:notify_time" json:"notify_time"`
	NotifyBeforeTime time.Time         `gorm:"column:notify_before_time" json:"notify_before_time"` // 单位秒
	TaskStatus       common.TaskStatus `gorm:"column:task_status;type:tinyint" json:"task_status"`
	StrategyData     string            `gorm:"column:strategy_data" json:"strategy_data"` // 策略数据
}

// TableName 指定表名
func (*TaskModel) TableName() string {
	return "tasks"
}

func ToDbModel(task *entity.Task) *TaskModel {
	return &TaskModel{
		TaskID:           task.TaskID,
		TaskName:         task.TaskName,
		TaskDesc:         task.TaskDesc,
		CreateTime:       task.CreateTime,
		TaskType:         int(task.TaskType),
		NotifyType:       task.NotifyType,
		NotifyTime:       task.NotifyTime,
		NotifyBeforeTime: task.NotifyBeforeTime,
		TaskStatus:       task.TaskStatus,
	}
}

// ToBizModel 转换为业务模型（假设需要隐藏某些字段）
func (t *TaskModel) ToBizModel() *entity.Task {
	return &entity.Task{
		TaskID:           t.TaskID,
		TaskName:         t.TaskName,
		TaskDesc:         t.TaskDesc,
		TaskType:         t.TaskType,
		TaskStatus:       t.TaskStatus,
		NotifyType:       t.NotifyType,
		NotifyTime:       t.NotifyTime,
		CreateTime:       t.CreateTime,
		NotifyBeforeTime: t.NotifyBeforeTime,
	}
}

func NewTaskRepository(db *gorm.DB) repository.TaskRepository {
	return &TaskRepositoryImpl{db: db}
}

func (r *TaskRepositoryImpl) GetTask(ctx context.Context, userId int64, taskId int64) (*entity.Task, error) { //-
	taskModel := &TaskModel{}

	err := r.db.Where("task_id = ? and user_id = ?", taskId, userId).First(task).Error
	if err != nil {
		logs.Errorf(ctx, "get task failed, err:%v", err)
		return nil, err
	}

	return taskModel.ToBizModel(), nil
}

func (r *TaskRepositoryImpl) UpdateTask(ctx context.Context, userId int64, task *entity.Task) error {
	taskModel := ToDbModel(task)
	err := r.db.Model(task).Where("task_id = ? and user_id = ?", task.TaskID, userId).Updates(taskModel).Error
	if err != nil {
		logs.Errorf(ctx, "update task failed, err:%v", err)
		return err
	}
	return nil
}

func (r *TaskRepositoryImpl) CreateTask(ctx context.Context, userId int64, task *entity.Task) error {

	taskModel := ToDbModel(task)
	err := r.db.Create(taskModel).Error
	if err != nil {
		logs.Errorf(ctx, "create task failed, err:%v", err)
		return err
	}
	return nil
}

func (r *TaskRepositoryImpl) DeleteTask(ctx context.Context, userId int64, taskId int64) error {
	err := r.db.Where("task_id = ? and user_id = ?", taskId, userId).Delete(&TaskModel{}).Error
	if err != nil {
		logs.Errorf(ctx, "delete task failed, err:%v", err)
		return err
	}
	return nil
}

func (r *TaskRepositoryImpl) GetTaskList(ctx context.Context, userId int64, offset, size int32) ([]*entity.Task, error) {

	taskModels := make([]*TaskModel, 0)
	err := r.db.Where("user_id = ?", userId).Offset(int(offset)).Limit(int(size)).Find(&taskModels).Error
	if err != nil {
		logs.Errorf(ctx, "get task list failed, err:%v", err)
		return nil, err
	}

	tasks := make([]*entity.Task, 0, len(taskModels))
	for _, taskModel := range taskModels {
		tasks = append(tasks, taskModel.ToBizModel())
	}
	return tasks, nil
}
