package service

import (
	"context"
	"notify/internal/common"
	"notify/internal/domain/entity"
	"notify/internal/domain/repository"
	"notify/internal/infra/notify"

	"github.com/sherlockhua/koala/logs"
)

type TaskService interface {
	GetTask(ctx context.Context, userId int64, taskId int64) (*entity.Task, error)
	UpdateTask(ctx context.Context, userId int64, task *entity.Task) error
	CreateTask(ctx context.Context, userId int64, task *entity.Task) error
	DeleteTask(ctx context.Context, userId int64, taskId int64) error
	TriggerTask(ctx context.Context) error
	CreateTemplate(ctx context.Context, template *entity.TaskTemplate) error
	GetTemplate(ctx context.Context, templateId int64) (*entity.TaskTemplate, error)
	GetTemplateList(ctx context.Context, offset, size int32) ([]*entity.TaskTemplate, error)
	UpdateTemplate(ctx context.Context, template *entity.TaskTemplate) error
	DeleteTemplate(ctx context.Context, templateId int64) error
}

type taskServiceImp struct {
	taskRepo       repository.TaskRepository
	templateRepo   repository.TemplateRepository
	accountService AccountService
	notifyService  notify.Notify
}

func NewTaskService(taskRepo repository.TaskRepository,
	service AccountService, notifyService notify.Notify, templateRepo repository.TemplateRepository) TaskService {
	return &taskServiceImp{
		taskRepo:       taskRepo,
		accountService: service,
		notifyService:  notifyService,
		templateRepo:   templateRepo,
	}
}

func (s *taskServiceImp) GetTask(ctx context.Context, userId int64, taskId int64) (*entity.Task, error) {
	logs.Debugf(ctx, "getting task, userId:%d, taskId:%d", userId, taskId)
	return s.taskRepo.GetTask(ctx, userId, taskId)
}

func (s *taskServiceImp) UpdateTask(ctx context.Context, userId int64, task *entity.Task) error {
	logs.Debugf(ctx, "updating task, userId:%d, task:%+v", userId, task)
	return s.taskRepo.UpdateTask(ctx, userId, task)
}

func (s *taskServiceImp) CreateTask(ctx context.Context, userId int64, task *entity.Task) error {
	logs.Debugf(ctx, "creating task, userId:%d, task:%+v", userId, task)
	hasBalane, err := s.accountService.HasBalance(ctx, userId)
	if err != nil {
		logs.Errorf(ctx, "check account balance failed, err:%v, user_id:%v", err, userId)
		return err
	}
	if !hasBalane {
		logs.Errorf(ctx, "account balance not enough, user_id:%v", userId)
		return common.ErrAccountBalanceNotEnough
	}

	task.TaskStatus = common.TaskStatusRunning
	return s.taskRepo.CreateTask(ctx, userId, task)
}

func (s *taskServiceImp) DeleteTask(ctx context.Context, userId int64, taskId int64) error {
	logs.Debugf(ctx, "deleting task, userId:%d, taskId:%d", userId, taskId)
	return s.taskRepo.DeleteTask(ctx, userId, taskId)
}

func (s *taskServiceImp) TriggerTask(ctx context.Context) error {
	logs.Debugf(ctx, "triggering task")
	tasks, err := s.taskRepo.GetAllTaskList(ctx, 0, 100, nil,
		[]common.TaskStatus{common.TaskStatusRunning})
	if err != nil {
		logs.Errorf(ctx, "get task list failed, err:%v", err)
		return err
	}
	for _, task := range tasks {
		err = s.processOneTask(ctx, task)
		if err != nil {
			logs.Errorf(ctx, "process task failed, err:%v, task:%+v", err, task)
			continue
		}
	}

	return nil
}

func (s *taskServiceImp) processOneTask(ctx context.Context, task *entity.Task) error {

	switch task.TaskStatus {
	case common.TaskStatusRunning:
		return s.processRunningTask(ctx, task)
	default:
		return nil
	}

}

func (s *taskServiceImp) processRunningTask(ctx context.Context, task *entity.Task) error {
	// 1. 检查是否到达通知时间
	notifyTimeResult := task.Strategy.IsTimeToNotify(ctx)
	switch notifyTimeResult {
	case common.NotifyTimeResultTimeReady:
		return s.processTimeReadyTask(ctx, task)
	case common.NotifyTimeResultBeforeTimeReady:
		return s.processBeforeTimeReadyTask(ctx, task)
	default:
		return nil
	}
}

func (s *taskServiceImp) processTimeReadyTask(ctx context.Context, task *entity.Task) error {
	return nil
}

func (s *taskServiceImp) processBeforeTimeReadyTask(ctx context.Context, task *entity.Task) error {
	return nil
}

func (s *taskServiceImp) CreateTemplate(ctx context.Context, template *entity.TaskTemplate) error {
	logs.Debugf(ctx, "creating template, template:%+v", template)
	return s.templateRepo.CreateTemplate(ctx, template)
}

func (s *taskServiceImp) GetTemplate(ctx context.Context, templateId int64) (*entity.TaskTemplate, error) {
	logs.Debugf(ctx, "getting template, templateId:%d", templateId)
	return s.templateRepo.GetTemplate(ctx, templateId)
}

func (s *taskServiceImp) GetTemplateList(ctx context.Context, offset, size int32) ([]*entity.TaskTemplate, error) {
	logs.Debugf(ctx, "getting template list, offset:%d, size:%d", offset, size)
	return s.templateRepo.GetTemplateList(ctx, offset, size, nil)
}

func (s *taskServiceImp) UpdateTemplate(ctx context.Context, template *entity.TaskTemplate) error {
	logs.Debugf(ctx, "updating template, template:%+v", template)
	return s.templateRepo.UpdateTemplate(ctx, template)
}

func (s *taskServiceImp) DeleteTemplate(ctx context.Context, templateId int64) error {
	logs.Debugf(ctx, "deleting template, templateId:%d", templateId)
	return s.templateRepo.DeleteTemplate(ctx, templateId)
}
