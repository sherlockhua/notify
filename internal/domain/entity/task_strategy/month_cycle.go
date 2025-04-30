package task_strategy

import (
	"context"
	"encoding/json"
	"fmt"
	"notify/internal/common"
	"time"
)

type MonthCycleStrategy struct {
	IntervalMonths      int      `json:"interval_months"`
	Monthdays           []int    `json:"month_days"`
	StartNotifyTimes    []string `json:"start_notify_times"`
	StartNotifyMonth    int      `json:"start_notify_Month"`
	NotifyBeforeSeconds int      `json:"notify_before_seconds"`
	timeUtil            common.TimeUtil
}

func NewMonthCycleStrategy(data string, timeUtil common.TimeUtil) (TaskStrategy, error) {
	inst := &MonthCycleStrategy{
		timeUtil: timeUtil,
	}
	err := json.Unmarshal([]byte(data), inst)
	if err != nil {
		return inst, err
	}

	return inst, nil
}
func (s *MonthCycleStrategy) match(monthDays []int, monthDay int, thisMonth time.Month) bool {
	if (int(thisMonth)-int(s.StartNotifyMonth))%s.IntervalMonths != 0 {
		return false
	}
	for _, s := range monthDays {
		if s == monthDay {
			return true
		}
	}
	return false
}

func (s *MonthCycleStrategy) IsTimeToNotify(ctx context.Context) common.NotifyTimeResult {
	currentTime := s.timeUtil.GetCurrentTime().UTC()
	monthDay := currentTime.Day()
	thisMonth := currentTime.Month()

	if !s.match(s.Monthdays, monthDay, thisMonth) {
		return common.NotifyTimeResultTimeNotReady
	}

	for _, startTime := range s.StartNotifyTimes {
		dateStr := fmt.Sprintf("%02d-%02d-%02d %s", currentTime.Year(), currentTime.Month(), currentTime.Day(), startTime)
		t, err := time.Parse(common.TimeLayoutWithoutZone, dateStr)
		if err != nil {
			return common.NotifyTimeResultTimeNotReady
		}

		// 计算当前时间和任务的通知时间
		days := common.DaysBetween(t, currentTime)
		expectedNotifyTime := t.Add(time.Duration(days*7*24*3600) * time.Second)
		notifyBeforeTime := expectedNotifyTime.Add(-1 * time.Duration(s.NotifyBeforeSeconds) * time.Second)

		if currentTime.Before(notifyBeforeTime) {
			continue
		}

		if currentTime.Sub(notifyBeforeTime).Seconds() <= TaskTriggerDiffSeconds {
			return common.NotifyTimeResultBeforeTimeReady
		}

		if currentTime.Sub(expectedNotifyTime).Seconds() <= TaskTriggerDiffSeconds {
			return common.NotifyTimeResultTimeReady
		}
	}

	return common.NotifyTimeResultTimeNotReady
}
