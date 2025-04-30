package repository

import (
	"context"
	"database/sql"

	"notify/internal/domain/entity"
)

type SubscriptionRepository struct {
	db *sql.DB
}

func NewSubscriptionRepository(db *sql.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) Save(ctx context.Context, sub *entity.Subscription) error {
	// 实现数据库保存逻辑
	return nil
}

func (r *SubscriptionRepository) FindByID(ctx context.Context, subscriptID int64) (*entity.Subscription, error) {
	// 实现数据库查询逻辑
	return nil, nil
}
