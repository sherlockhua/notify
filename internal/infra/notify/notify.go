package notify

import (
	"context"
	"fmt"
	"math/rand"
)

type Notify interface {
	NotifyTelephone
	NotifySms
}

type NotifyTelephone interface {
	SendTelephone(ctx context.Context, userId int64, taskId int64) error
}
type NotifySms interface {
	SendSms(ctx context.Context, phoneNumber, templateCode, templateParams string) error
}

type NotifyImpl struct {
	SmsServiceList       []NotifySms
	TelephoneServiceList []NotifyTelephone
}

type SmsChannel int

const (
	SmsTypeAliyun  SmsChannel = iota
	SmsTypeTencent SmsChannel = iota
)

type PhoneChannel int

const (
	PhoneChannelAliyun  PhoneChannel = iota
	PhoneChannelTencent PhoneChannel = iota
)

type SmsConf struct {
	AccessKeyID     string
	AccessKeySecret string
	SignName        string
	SmsType         SmsChannel
}
type NotifySmsConf struct {
	SmsConfList []SmsConf
}

type PhoneConf struct {
	AccessKeyID     string
	AccessKeySecret string
	SignName        string
	PhoneType       PhoneChannel
}

type NotifyPhoneConf struct {
	PhoneConfList []PhoneConf
}

func NewNotifyImpl(smsConf NotifySmsConf, phoneConf NotifyPhoneConf) (Notify, error) {
	inst := &NotifyImpl{}
	for _, conf := range smsConf.SmsConfList {
		switch conf.SmsType {
		case SmsTypeAliyun:
			sms, err := NewSMSAliyun(conf.AccessKeyID, conf.AccessKeySecret, conf.SignName)
			if err != nil {
				return nil, err
			}
			inst.SmsServiceList = append(inst.SmsServiceList, sms)
		case SmsTypeTencent:
			return nil, fmt.Errorf("not support")
		}
	}

	for _, conf := range phoneConf.PhoneConfList {
		switch conf.PhoneType {
		case PhoneChannelAliyun:
			return nil, fmt.Errorf("not support")
		case PhoneChannelTencent:
			return nil, fmt.Errorf("not support")
		}
	}

	return inst, nil
}

func (n *NotifyImpl) SendSms(ctx context.Context, phoneNumber, templateCode, templateParams string) error {
	//随机选择一个sms服务，调度算法后续再优化
	idx := rand.Intn(len(n.SmsServiceList))
	sms := n.SmsServiceList[idx]
	err := sms.SendSms(ctx, phoneNumber, templateCode, templateParams)
	if err != nil {
		return err
	}
	return nil
}

func (n *NotifyImpl) SendTelephone(ctx context.Context, userId int64, taskId int64) error {
	return nil
}
