package application

import (
	"context"
	"github.com/sherlockhua/koala/logs"
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

	err := r.taskService.TriggerTask(ctx)
	if err != nil {
		logs.Errorf(ctx, "TriggerTask failed, err:%v", err)
		return
	}

	logs.Infof(ctx, "TriggerTask success")
}
