package common

import "errors"

var (
	ErrAccountBalanceNotEnough = errors.New("account balance not enough")
	ErrAccountNotFound         = errors.New("account not found")
	ErrAccountStatusDisable    = errors.New("account status disable")
	ErrInvalidTaskType         = errors.New("invalid task type")
	ErrInvalidColumnName       = errors.New("invalid column name")
	ErrSendSmsFailed           = errors.New("send sms failed")
)
