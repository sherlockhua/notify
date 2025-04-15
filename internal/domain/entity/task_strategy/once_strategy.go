package task_strategy

import (
	"context"
	"encoding/json"
	"notify/internal/common"
	"time"
)

type OnceStrategy struct {
	StartNotifyTime     time.Time `json:"start_notify_time"`
	NotifyBeforeSeconds int       `json:"notify_before_seconds"`
	timeUtil            common.TimeUtil
}

func NewOnceStrategy(data string, timeUtil common.TimeUtil) (TaskStrategy, error) {
	inst := &OnceStrategy{
		timeUtil: timeUtil,
	}

	err := json.Unmarshal([]byte(data), inst)
	if err != nil {
		return inst, err
	}

	return inst, nil
}
func (s *OnceStrategy) IsTimeToNotify(ctx context.Context) common.NotifyTimeResult {
	// 计算当前时间和任务的通知时间
	currentTime := time.Now().UTC()
	notifyBeforeTime := s.StartNotifyTime.Add(-1 * time.Duration(s.NotifyBeforeSeconds) * time.Second)
	if currentTime.Before(notifyBeforeTime) {
		return common.NotifyTimeResultTimeNotReady
	}

	if currentTime.After(notifyBeforeTime) && currentTime.Before(s.StartNotifyTime) {
		return common.NotifyTimeResultBeforeTimeReady
	}
	return common.NotifyTimeResultTimeReady
}
