package entity

import (
	"notify/internal/common"
	"time"
)

type Account struct {
	AccountID      int64                `json:"account_id"`
	AccountName    string               `json:"account_name"`
	AccountDesc    string               `json:"account_desc"`
	CreateTime     time.Time            `json:"create_time"`
	AccountType    int                  `json:"account_type"`
	AccountStatus  common.AccountStatus `json:"account_status"`
	AccountBalance Money                `json:"account_balance"`
	UserID         int64                `json:"user_id"`
}
