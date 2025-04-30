package task_strategy

import (
	"context"
	"encoding/json"
	"fmt"
	"notify/internal/common"
	"time"
)

type YearCycleStrategy struct {
	IntervalYears       int      `json:"interval_Years"`
	Months              []int    `json:"months"`
	Days                []int    `json:"days"`
	StartNotifyTimes    []string `json:"start_notify_times"`
	StartNotifyYear     int      `json:"start_notify_Year"`
	NotifyBeforeSeconds int      `json:"notify_before_seconds"`
	timeUtil            common.TimeUtil
}

func NewYearCycleStrategy(data string, timeUtil common.TimeUtil) (TaskStrategy, error) {
	inst := &YearCycleStrategy{
		timeUtil: timeUtil,
	}
	err := json.Unmarshal([]byte(data), inst)
	if err != nil {
		return inst, err
	}

	return inst, nil
}
func (s *YearCycleStrategy) match(months []int, days []int, month, day int, thisYear int) bool {
	if (thisYear-s.StartNotifyYear)%s.IntervalYears != 0 {
		return false
	}
	for _, s := range months {
		if s == month {
			for _, d := range days {
				if d == day {
					return true
				}
			}
		}
	}
	return false
}

func (s *YearCycleStrategy) IsTimeToNotify(ctx context.Context) common.NotifyTimeResult {
	currentTime := s.timeUtil.GetCurrentTime().UTC()
	day := currentTime.Day()
	month := int(currentTime.Month())
	thisYear := currentTime.Year()

	if !s.match(s.Months, s.Days, month, day, thisYear) {
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
