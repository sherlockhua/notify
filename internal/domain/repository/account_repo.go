package repository

import (
	"context"
	"notify/internal/domain/entity"
)

type AccountRepository interface {
	CreateAccount(ctx context.Context, userId int64, account *entity.Account) error
	UpdateAccount(ctx context.Context, userId int64, account *entity.Account) error
	GetAccount(ctx context.Context, userId int64) (*entity.Account, error)
}
