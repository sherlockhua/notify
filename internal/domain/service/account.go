package service

import (
	"context"
	"notify/internal/domain/entity"
	"notify/internal/domain/repository"
)

type AccountService interface {
	CreateAccount(ctx context.Context, userId int64, account *entity.Account) error
	UpdateAccount(ctx context.Context, userId int64, account *entity.Account) error
	GetAccount(ctx context.Context, userId int64) (*entity.Account, error)
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
