package application

import (
	"context"
	"github.com/sherlockhua/koala/logs"
	"notify/internal/common"
	"notify/internal/domain/service"
	"time"
)

type TaskRunner interface {
	Start(ctx context.Context)
}

type TaskRunnerImpl struct {
	ctx            context.Context
	taskService    service.TaskService
	accountService service.AccountService
}

func NewTaskRunner(ctx context.Context, taskService service.TaskService, accountService service.AccountService) TaskRunner {
	return &TaskRunnerImpl{
		ctx:            ctx,
		taskService:    taskService,
		accountService: accountService,
	}
}
func (r *TaskRunnerImpl) Start(ctx context.Context) {
	go func() {
		if e := recover(); e != nil {
			logs.Errorf(ctx, "start task failed, err:%v", e)
		}

		for {
			r.handler(ctx)
			time.Sleep(time.Second)
		}
	}()
}

func (r *TaskRunnerImpl) handler(ctx context.Context) {
	// 1. 获取所有的任务
	tasks, err := r.taskService.GetTaskList(ctx, 0, 100)
	if err != nil {
		logs.Errorf(ctx, "get task list failed, err:%v", err)
		return
	}

	for _, task := range tasks {
		if task.TaskStatus == common.TaskStatusNotStart {
			logs.Infof(ctx, "task is not start, task_id:%v", task.TaskID)
			continue
		}
		if task.TaskStatus == common.TaskStatusRunning {
			logs.Infof(ctx, "task is running, task_id:%v", task.TaskID)
			continue
		}
		if task.TaskStatus == common.TaskStatusFinish {
			logs.Infof(ctx, "task is finish, task_id:%v", task.TaskID)
			continue
		}
		if task.TaskStatus == common.TaskStatusError {
			logs.Infof(ctx, "task is error, task_id:%v", task.TaskID)
			continue
		}
	}
}
