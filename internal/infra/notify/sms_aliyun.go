package notify

import (
	"context"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"notify/internal/common"
)

type SMSAliyun struct {
	accessKeyID     string
	accessKeySecret string
	signName        string
	client          *dysmsapi.Client
}

func NewSMSAliyun(accessKeyID, accessKeySecret, signName string) (NotifySms, error) {
	inst := &SMSAliyun{
		accessKeyID:     accessKeyID,
		accessKeySecret: accessKeySecret,
		signName:        signName,
	}

	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", accessKeyID, accessKeySecret)
	if err != nil {
		return nil, err
	}
	inst.client = client
	return inst, nil
}

func (s *SMSAliyun) SendSms(ctx context.Context, phoneNumber, templateCode, templateParams string) error {
	// 创建发送短信请求
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = s.signName
	request.TemplateCode = templateCode
	request.PhoneNumbers = phoneNumber
	request.TemplateParam = templateParams

	// 发送请求
	response, err := s.client.SendSms(request)
	if err != nil {
		return err
	}

	if response.Code != "OK" {
		return common.ErrSendSmsFailed
	}
	return nil
}
