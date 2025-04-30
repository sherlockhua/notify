package application

import (
	"context"
	"github.com/sherlockhua/koala/logs"
	"notify/internal/domain/entity"
	"notify/internal/domain/service"
)

type Appalition interface {
	CreateTask(ctx context.Context, userId int64, task *entity.Task) error
	Start(ctx context.Context)
}

type AppalitionImpl struct {
	ctx            context.Context
	taskService    service.TaskService
	accountService service.AccountService
	taskRunnerList []TaskRunner
}

func NewAppalition(ctx context.Context, taskService service.TaskService, accountService service.AccountService) Appalition {
	inst := &AppalitionImpl{
		ctx:         ctx,
		taskService: taskService,
	}
	inst.taskRunnerList = append(inst.taskRunnerList, NewTaskRunner(ctx, taskService, accountService))
	return inst
}

func (a *AppalitionImpl) CreateTask(ctx context.Context, userId int64, task *entity.Task) error {
	err := a.taskService.CreateTask(a.ctx, userId, task)
	if err != nil {
		logs.Errorf(ctx, "create task failed, err:%v, user_id:%v", err, userId)
		return err
	}
	return nil
}

func (a *AppalitionImpl) Start(ctx context.Context) {

	go func() {
		if e := recover(); e != nil {
			logs.Errorf(ctx, "start task failed, err:%v", e)
		}
		for _, runner := range a.taskRunnerList {
			runner.Start(ctx)
		}
	}()

}
