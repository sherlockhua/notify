package entity

import (
	"notify/internal/common"
	"time"
)

type Task struct {
	TaskID     int64     `json:"task_id"`
	TaskName   string    `json:"task_name"`
	TaskDesc   string    `json:"task_desc"`
	CreateTime time.Time `json:"create_time"`
	TaskType   int       `json:"task_type"`
	NotifyType int       `json:"notify_type"`
	NotifyTime time.Time `json:"notify_time"`
	//单位是秒
	NotifyBeforeTime int               `json:"notify_before_time"`
	TaskStatus       common.TaskStatus `json:"task_status"`
}
