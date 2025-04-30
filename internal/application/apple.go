package application

import (
	"context"
	"log"

	"notify/internal/domain/service"
)

type SubscriptionApplication struct {
	verificationSvc  *service.VerificationService
	subscriptionRepo model.SubscriptionRepository
	eventPublisher   EventPublisher
}

func NewSubscriptionApplication(
	verificationSvc *service.VerificationService,
	subscriptionRepo model.SubscriptionRepository,
	eventPublisher EventPublisher,
) *SubscriptionApplication {
	return &SubscriptionApplication{
		verificationSvc:  verificationSvc,
		subscriptionRepo: subscriptionRepo,
		eventPublisher:   eventPublisher,
	}
}

func (a *SubscriptionApplication) ProcessNewSubscription(
	ctx context.Context,
	userID string,
	receiptData string,
) (*model.Subscription, error) {
	// 1. 验证并创建订阅
	subscription, err := a.verificationSvc.VerifyAndCreateSubscription(ctx, userID, receiptData)
	if err != nil {
		return nil, err
	}

	// 2. 保存订阅
	if err := a.subscriptionRepo.Save(ctx, subscription); err != nil {
		return nil, err
	}

	// 3. 发布领域事件
	if err := a.eventPublisher.PublishSubscriptionCreated(ctx, subscription); err != nil {
		// 记录日志但继续流程
		log.Printf("failed to publish event: %v", err)
	}

	return subscription, nil
}
