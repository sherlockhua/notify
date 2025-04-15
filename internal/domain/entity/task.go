package entity

import (
	"notify/internal/common"
	"time"
)

type Task struct {
	TaskID          int64                `json:"task_id"`
	TaskName        string               `json:"task_name"`
	TaskDesc        string               `json:"task_desc"`
	CreateTime      time.Time            `json:"create_time"`
	UpdateTime      time.Time            `json:"update_time"`
	TaskType        common.TaskType      `json:"task_type"`
	NotifyChannel   common.NotifyChannel `json:"notify_channel"`
	StartNotifyTime time.Time            `json:"start_notify_time"`
	//单位是秒
	NotifyBeforeSeconds int               `json:"notify_before_seconds"`
	TaskStatus          common.TaskStatus `json:"task_status"`
	StrategyData        string            `json:"strategy_data"`
}
