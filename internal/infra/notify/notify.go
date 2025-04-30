package notify

import "context"

type Notify interface {
	SendSms(ctx context.Context, userId int64, taskId int64) error
	SendEmail(ctx context.Context, userId int64, taskId int64) error
	SendTelephone(ctx context.Context, userId int64, taskId int64) error
}

type NotifyImpl struct {
}

func NewNotify() Notify {
	return &NotifyImpl{}
}

func (n *NotifyImpl) SendSms(ctx context.Context, userId int64, taskId int64) error {
	return nil
}

func (n *NotifyImpl) SendEmail(ctx context.Context, userId int64, taskId int64) error {
	return nil
}

func (n *NotifyImpl) SendTelephone(ctx context.Context, userId int64, taskId int64) error {
	return nil
}
