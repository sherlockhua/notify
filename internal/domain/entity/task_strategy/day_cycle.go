package task_strategy

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sherlockhua/koala/logs"
	"notify/internal/common"
	"time"
)

type DayCycleStrategy struct {
	IntervalDays        int      `json:"interval_days"`
	StartNotifyTimes    []string `json:"start_notify_times"`
	StartNotifyDate     string   `json:"start_notify_date"`
	NotifyBeforeSeconds int      `json:"notify_before_seconds"`
	timeUtil            common.TimeUtil
}

func NewDayCycleStrategy(param string, util common.TimeUtil) (TaskStrategy, error) {
	inst := &DayCycleStrategy{
		timeUtil: util,
	}

	err := json.Unmarshal([]byte(param), inst)
	if err != nil {
		return nil, err
	}

	return inst, nil
}

func (d *DayCycleStrategy) IsTimeToNotify(ctx context.Context) common.NotifyTimeResult {
	// 计算当前时间和任务的通知时间
	currentTime := d.timeUtil.GetCurrentTime().UTC()
	for _, startTime := range d.StartNotifyTimes {
		t, err := time.Parse(common.TimeLayoutWithoutZone, fmt.Sprintf("%s %s", d.StartNotifyDate, startTime))
		if err != nil {
			logs.Errorf(ctx, "parse start notify time failed, err:%v, data:%v", err, d)
			return common.NotifyTimeResultTimeNotReady
		}

		days := common.DaysBetween(t, currentTime)
		if days%d.IntervalDays != 0 {
			return common.NotifyTimeResultTimeNotReady
		}

		expectedNotifyTime := t.Add(time.Duration(days*24*3600) * time.Second)
		notifyBeforeTime := expectedNotifyTime.Add(-1 * time.Duration(d.NotifyBeforeSeconds) * time.Second)

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
