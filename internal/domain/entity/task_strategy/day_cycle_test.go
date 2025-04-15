package task_strategy

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"time"

	"notify/internal/common"
	"testing"
)

func TestDayCycleStrategy_IsTimeToNotify(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言所有预期调用都发生了

	type args struct {
		data           string
		currentTime    time.Time
		currentTimeStr string
		result         common.NotifyTimeResult
	}

	tests := []args{
		{
			data: `{
				"interval_days": 1,
				"start_notify_times": ["14:00:00"],
				"start_notify_date": "2025-04-14",
				"notify_before_seconds": 300
			}`,
			currentTimeStr: "2025-04-15 14:00:00 +0000 UTC",
			result:         common.NotifyTimeResultTimeReady,
		},
		{
			data: `{
				"interval_days": 1,
				"start_notify_times": ["14:00:00"],
				"start_notify_date": "2025-04-14",
				"notify_before_seconds": 300
			}`,
			currentTimeStr: "2025-04-15 13:55:00 +0000 UTC",
			result:         common.NotifyTimeResultBeforeTimeReady,
		},
		{
			data: `{
				"interval_days": 1,
				"start_notify_times": ["14:00:00"],
				"start_notify_date": "2025-04-14",
				"notify_before_seconds": 300
			}`,
			currentTimeStr: "2025-04-15 13:50:00 +0000 UTC",
			result:         common.NotifyTimeResultTimeNotReady,
		},
		{
			data: `{
				"interval_days": 2,
				"start_notify_times": ["14:00:00"],
				"start_notify_date": "2025-04-14",
				"notify_before_seconds": 300
			}`,
			currentTimeStr: "2025-04-16 14:00:00 +0000 UTC",
			result:         common.NotifyTimeResultTimeReady,
		},
		{
			data: `{
				"interval_days": 2,
				"start_notify_times": ["14:00:00"],
				"start_notify_date": "2025-04-14",
				"notify_before_seconds": 300
			}`,
			currentTimeStr: "2025-04-16 13:55:00 +0000 UTC",
			result:         common.NotifyTimeResultBeforeTimeReady,
		},
		{
			data: `{
				"interval_days": 2,
				"start_notify_times": ["14:00:00"],
				"start_notify_date": "2025-04-14",
				"notify_before_seconds": 300
			}`,
			currentTimeStr: "2025-04-15 14:00:00 +0000 UTC",
			result:         common.NotifyTimeResultTimeNotReady,
		},
		{
			data: `{
				"interval_days": 2,
				"start_notify_times": ["8:00:00","12:00:00", "18:00:00"],
				"start_notify_date": "2025-04-14",
				"notify_before_seconds": 300
			}`,
			currentTimeStr: "2025-04-16 8:00:00 +0000 UTC",
			result:         common.NotifyTimeResultTimeReady,
		},
		{
			data: `{
				"interval_days": 2,
				"start_notify_times": ["8:00:00","12:00:00", "18:00:00"],
				"start_notify_date": "2025-04-14",
				"notify_before_seconds": 300
			}`,
			currentTimeStr: "2025-04-16 12:00:00 +0000 UTC",
			result:         common.NotifyTimeResultTimeReady,
		},
		{
			data: `{
				"interval_days": 2,
				"start_notify_times": ["8:00:00","12:00:00", "18:00:00"],
				"start_notify_date": "2025-04-14",
				"notify_before_seconds": 300
			}`,
			currentTimeStr: "2025-04-16 18:00:00 +0000 UTC",
			result:         common.NotifyTimeResultTimeReady,
		},
	}

	for _, one := range tests {

		one.currentTime, _ = time.Parse("2006-01-02 15:04:05 -0700 MST", one.currentTimeStr)
		t.Logf("test data: %v current_time:%v", one.data, one.currentTime)

		timeUtil := common.NewMockTimeUtil(ctrl)
		timeUtil.EXPECT().GetCurrentTime().Return(one.currentTime).AnyTimes()
		dayCycleStrategy, err := NewDayCycleStrategy(one.data, timeUtil)
		assert.Equal(t, nil, err)
		result := dayCycleStrategy.IsTimeToNotify(context.Background())
		assert.Equal(t, one.result, result, one)
	}
}
