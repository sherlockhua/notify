package repository

import (
	"context"
	"notify/internal/common"
	"notify/internal/domain/entity"
	"notify/internal/domain/entity/task_strategy"
	"notify/internal/domain/repository"
	"time"

	"github.com/sherlockhua/koala/logs"
	"gorm.io/gorm"
)

type TaskRepositoryImpl struct {
	db              *gorm.DB //+
	strategyFactory task_strategy.TaskStrategyFactory
}

type TaskModel struct {
	ID           int64             `gorm:"column:id" json:"id"`
	TaskID       int64             `gorm:"column:task_id" json:"task_id"`
	UserID       int64             `gorm:"column:user_id" json:"user_id"`
	TaskName     string            `gorm:"column:task_name" json:"task_name"`
	TaskDesc     string            `gorm:"column:task_desc" json:"task_desc"`
	CreateTime   time.Time         `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"`
	TaskType     int               `gorm:"column:task_type" json:"task_type"`
	TaskStatus   common.TaskStatus `gorm:"column:task_status;type:tinyint" json:"task_status"`
	StrategyData string            `gorm:"column:strategy_data" json:"strategy_data"` // 策略数据
}

// TableName 指定表名
func (*TaskModel) TableName() string {
	return "tasks"
}

func ToDbModel(task *entity.Task) *TaskModel {
	return &TaskModel{
		TaskID:       task.TaskID,
		UserID:       task.UserID,
		TaskName:     task.TaskName,
		TaskDesc:     task.TaskDesc,
		CreateTime:   task.CreateTime,
		TaskType:     int(task.TaskType),
		TaskStatus:   task.TaskStatus,
		StrategyData: task.StrategyData,
	}
}

// ToBizModel 转换为业务模型（假设需要隐藏某些字段）
func (t *TaskModel) ToBizModel(ctx context.Context, strategyFactory task_strategy.TaskStrategyFactory) (*entity.Task, error) {
	task := &entity.Task{
		TaskID:     t.TaskID,
		UserID:     t.UserID,
		TaskName:   t.TaskName,
		TaskDesc:   t.TaskDesc,
		TaskType:   common.TaskType(t.TaskType),
		TaskStatus: t.TaskStatus,
		CreateTime: t.CreateTime,
	}

	var err error
	task.Strategy, err = strategyFactory.CreateTaskStrategy(ctx, task.TaskType, t.StrategyData)
	if err != nil {
		logs.Errorf(ctx, "create task strategy failed, err:%v", err)
		return nil, err
	}
	return task, nil
}

func NewTaskRepository(db *gorm.DB, factory task_strategy.TaskStrategyFactory) repository.TaskRepository {
	return &TaskRepositoryImpl{db: db, strategyFactory: factory}
}

func (r *TaskRepositoryImpl) GetTask(ctx context.Context, userId int64, taskId int64) (*entity.Task, error) { //-
	taskModel := &TaskModel{}
	err := r.db.Where("task_id = ? and user_id = ?", taskId, userId).First(taskModel).Error
	if err != nil {
		logs.Errorf(ctx, "get task failed, err:%v", err)
		return nil, err
	}

	return taskModel.ToBizModel(ctx, r.strategyFactory)
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

func (r *TaskRepositoryImpl) GetTaskList(ctx context.Context, userId int64, offset, size int32,
	conditions map[string]interface{}, statusConditions []common.TaskStatus) ([]*entity.Task, error) {

	taskModels := make([]*TaskModel, 0)
	whereSql := r.db.Where("user_id = ?", userId)
	if len(conditions) > 0 {
		validConditions := make(map[string]interface{})
		for k, v := range conditions {
			isValid, err := common.IsValidColumnName(TaskModel{}, k)
			if err != nil {
				logs.Errorf(ctx, "check column name failed, err:%v, conditions:%v", err, conditions)
				return nil, err
			}
			if isValid {
				validConditions[k] = v
			} else {
				logs.Errorf(ctx, "invalid column name: %s, conditions:%v", k, conditions)
				return nil, common.ErrInvalidColumnName
			}
		}
		whereSql = whereSql.Where(validConditions)
	}

	if len(statusConditions) > 0 {
		whereSql = whereSql.Where("task_status in (?)", statusConditions)
	}
	err := whereSql.Offset(int(offset)).Limit(int(size)).Find(&taskModels).Error
	if err != nil {
		logs.Errorf(ctx, "get task list failed, err:%v", err)
		return nil, err
	}

	tasks := make([]*entity.Task, 0, len(taskModels))
	for _, taskModel := range taskModels {
		task, err := taskModel.ToBizModel(ctx, r.strategyFactory)
		if err != nil {
			logs.Errorf(ctx, "convert task model to biz model failed, err:%v", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *TaskRepositoryImpl) GetAllTaskList(ctx context.Context, offset, size int32,
	conditions map[string]interface{}, statusConditions []common.TaskStatus) ([]*entity.Task, error) {
	taskModels := make([]*TaskModel, 0)
	whereSql := r.db
	if len(conditions) > 0 {
		validConditions := make(map[string]interface{})
		for k, v := range conditions {
			isValid, err := common.IsValidColumnName(TaskModel{}, k)
			if err != nil {
				logs.Errorf(ctx, "check column name failed, err:%v, conditions:%v", err, conditions)
				return nil, err
			}
			if isValid {
				validConditions[k] = v
			} else {
				logs.Errorf(ctx, "invalid column name: %s, conditions:%v", k, conditions)
				return nil, common.ErrInvalidColumnName
			}
		}
		whereSql = whereSql.Where(validConditions)
	}

	if len(statusConditions) > 0 {
		whereSql = whereSql.Where("task_status in (?)", statusConditions)
	}
	err := whereSql.Offset(int(offset)).Limit(int(size)).Find(&taskModels).Error
	if err != nil {
		logs.Errorf(ctx, "get task list failed, err:%v", err)
		return nil, err
	}

	tasks := make([]*entity.Task, 0, len(taskModels))
	for _, taskModel := range taskModels {
		task, err := taskModel.ToBizModel(ctx, r.strategyFactory)
		if err != nil {
			logs.Errorf(ctx, "convert task model to biz model failed, err:%v", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *TaskRepositoryImpl) CreateTaskTemplate(ctx context.Context, template *entity.TaskTemplate) error {
	return nil
}
