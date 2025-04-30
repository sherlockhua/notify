package service

import (
	"context"
	"github.com/sherlockhua/koala/logs"
	"notify/internal/common"
	"notify/internal/domain/entity"
	"notify/internal/domain/repository"
)

type AccountService interface {
	CreateAccount(ctx context.Context, userId int64, account *entity.Account) error
	UpdateAccount(ctx context.Context, userId int64, account *entity.Account) error
	GetAccount(ctx context.Context, userId int64) (*entity.Account, error)
	HasBalance(ctx context.Context, userId int64) (bool, error)
}

type accountServiceImp struct {
	accountRepo repository.AccountRepository
}

func NewAccountService(accountRepo repository.AccountRepository) AccountService {
	return &accountServiceImp{accountRepo: accountRepo}
}
func (s *accountServiceImp) CreateAccount(ctx context.Context, userId int64, account *entity.Account) error {
	return s.accountRepo.CreateAccount(ctx, userId, account)
}

func (s *accountServiceImp) UpdateAccount(ctx context.Context, userId int64, account *entity.Account) error {
	return s.accountRepo.UpdateAccount(ctx, userId, account)
}

func (s *accountServiceImp) GetAccount(ctx context.Context, userId int64) (*entity.Account, error) {
	return s.accountRepo.GetAccount(ctx, userId)
}

func (s *accountServiceImp) HasBalance(ctx context.Context, userId int64) (bool, error) {
	//先判断是否有余额
	account, err := s.GetAccount(ctx, userId)
	if err != nil {
		logs.Errorf(ctx, "get account failed, err:%v, user_id:%v", err, userId)
		return false, err
	}

	if account.AccountStatus == common.AccountStatusDisable {
		logs.Infof(ctx, "account status is disable, user_id:%v", userId)
		return false, common.ErrAccountStatusDisable
	}

	if account.AccountBalance.Amount <= 0 {
		logs.Infof(ctx, "account balance is not enough, user_id:%v, balance:%v", userId, account.AccountBalance.Amount)
		return false, nil
	}

	return true, nil
}
