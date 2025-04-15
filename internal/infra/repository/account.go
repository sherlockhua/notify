package repository

import (
	"context"
	"time"

	"github.com/sherlockhua/koala/logs"
	"gorm.io/gorm"
	"notify/internal/common"
	"notify/internal/domain/entity"
	"notify/internal/domain/repository"
)

type AccountRepositoryImpl struct {
	db *gorm.DB //+
}
type AccountModel struct {
	ID             int64                `gorm:"column:id" json:"id"`
	AccountID      int64                `gorm:"column:account_id" json:"account_id"`
	AccountName    string               `gorm:"column:account_name"  json:"account_name"`
	AccountDesc    string               `gorm:"column:account_desc" json:"account_desc"`
	CreateTime     time.Time            `gorm:"column:create_time" json:"create_time"`
	AccountType    int                  `gorm:"column:account_type" json:"account_type"`
	AccountStatus  common.AccountStatus `gorm:"column:account_status" json:"account_status"`
	AccountBalance int64                `gorm:"column:account_balance" json:"account_balance"`
	Currency       string               `gorm:"column:currency" json:"currency"`
	UserID         int64                `gorm:"column:user_id" json:"user_id"`
}

// TableName 指定表名
func (*AccountModel) TableName() string {
	return "account"
}

func ToAccountDbModel(account *entity.Account) *AccountModel {
	return &AccountModel{
		AccountID:      account.AccountID,
		AccountName:    account.AccountName,
		AccountDesc:    account.AccountDesc,
		CreateTime:     account.CreateTime,
		AccountType:    account.AccountType,
		AccountStatus:  account.AccountStatus,
		AccountBalance: account.AccountBalance.Amount,
		Currency:       account.AccountBalance.Currency,
		UserID:         account.UserID,
	}
}

// ToBizModel 转换为业务模型（假设需要隐藏某些字段）
func (t *AccountModel) ToBizModel() *entity.Account {
	return &entity.Account{
		AccountID:     t.AccountID,
		AccountName:   t.AccountName,
		AccountDesc:   t.AccountDesc,
		CreateTime:    t.CreateTime,
		AccountType:   t.AccountType,
		AccountStatus: t.AccountStatus,
		AccountBalance: entity.Money{
			Amount:   t.AccountBalance,
			Currency: t.Currency,
		},
		UserID: t.UserID,
	}
}

func NewAccountRepository(db *gorm.DB) repository.AccountRepository {
	return &AccountRepositoryImpl{db: db}
}

func (r *AccountRepositoryImpl) GetAccount(ctx context.Context, userId int64) (*entity.Account, error) { //-
	accountModel := &AccountModel{}
	err := r.db.Where("user_id =?", userId).First(accountModel).Error
	if err != nil {
		logs.Errorf(ctx, "get account failed, err:%v", err)
		return nil, err
	}
	return accountModel.ToBizModel(), nil
}
func (r *AccountRepositoryImpl) UpdateAccount(ctx context.Context, userId int64, account *entity.Account) error {
	accountModel := ToAccountDbModel(account)
	err := r.db.Model(account).Where("account_id =? and user_id =?", account.AccountID, userId).Updates(accountModel).Error
	if err != nil {
		logs.Errorf(ctx, "update account failed, err:%v", err)
		return err
	}
	return nil
}

func (r *AccountRepositoryImpl) CreateAccount(ctx context.Context, userId int64, account *entity.Account) error {
	accountModel := ToAccountDbModel(account)
	err := r.db.Create(accountModel).Error
	if err != nil {
		logs.Errorf(ctx, "create account failed, err:%v", err)
		return err
	}
	return nil
}
