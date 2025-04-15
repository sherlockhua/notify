package task_strategy

import (
	"context"
	"encoding/json"
	"fmt"
	"notify/internal/common"
	"time"
)

type WeekCycleStrategy struct {
	IntervalWeeks       int            `json:"interval_weeks"`
	Weekdays            []time.Weekday `json:"weekdays"`
	StartNotifyTimes    []string       `json:"start_notify_times"`
	StartNotifyWeek     int            `json:"start_notify_week"`
	NotifyBeforeSeconds int            `json:"notify_before_seconds"`
	timeUtil            common.TimeUtil
}

func NewWeekCycleStrategy(data string, timeUtil common.TimeUtil) (TaskStrategy, error) {
	inst := &WeekCycleStrategy{
		timeUtil: timeUtil,
	}
	err := json.Unmarshal([]byte(data), inst)
	if err != nil {
		return inst, err
	}

	return inst, nil
}
func (s *WeekCycleStrategy) match(slice []time.Weekday, item time.Weekday, thisWeek int) bool {
	if (thisWeek-s.StartNotifyWeek)%s.IntervalWeeks != 0 {
		return false
	}
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func (s *WeekCycleStrategy) IsTimeToNotify(ctx context.Context) common.NotifyTimeResult {
	currentTime := s.timeUtil.GetCurrentTime().UTC()
	weekDay := currentTime.Weekday()
	_, thisWeek := currentTime.ISOWeek()

	if !s.match(s.Weekdays, weekDay, thisWeek) {
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
		if days%s.IntervalWeeks != 0 {
			continue
		}

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
