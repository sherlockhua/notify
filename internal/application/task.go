package application

import (
	"context"
	"github.com/sherlockhua/koala/logs"
	"notify/internal/common"
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
}

func NewAppalition(ctx context.Context, taskService service.TaskService) Appalition {
	return &AppalitionImpl{
		ctx:         ctx,
		taskService: taskService,
	}
}

func (a *AppalitionImpl) CreateTask(ctx context.Context, userId int64, task *entity.Task) error {

	//先判断是否有余额
	account, err := a.accountService.GetAccount(ctx, userId)
	if err != nil {
		logs.Errorf(ctx, "get account failed, err:%v, user_id:%v", err, userId)
		return err
	}

	if account.AccountStatus == common.AccountStatusDisable {
		logs.Infof(ctx, "account status is disable, user_id:%v", userId)
		return common.ErrAccountStatusDisable
	}

	if account.AccountBalance.Amount <= 0 {
		logs.Infof(ctx, "account balance is not enough, user_id:%v, balance:%v", userId, account.AccountBalance.Amount)
		return common.ErrAccountBalanceNotEnough
	}
	err = a.taskService.CreateTask(a.ctx, userId, task)
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

	}()

}
