package common

type TaskStatus int

const (
	TaskStatusNotStart TaskStatus = 1
	TaskStatusRunning  TaskStatus = 2
	TaskStatusFinished TaskStatus = 3
	TaskStatusCannel   TaskStatus = 4
)

type TaskType int

const (
	TaskTypeOnce        TaskType = 1
	TaskTypeDayCycle    TaskType = 2
	TaskTypeWeekCycle   TaskType = 3
	TaskTypeMonthCycle  TaskType = 4
	TaskTypeYearCycle   TaskType = 5
	TaskTypeCustomCycle TaskType = 6
)

type NotifyChannel int

const (
	NotifyChannelPhone NotifyChannel = 1
	NotifyChannelSMS   NotifyChannel = 2
)

type NotifyTimeResult int

const (
	NotifyTimeResultBeforeTimeReady NotifyTimeResult = 1
	NotifyTimeResultTimeReady       NotifyTimeResult = 2
	NotifyTimeResultTimeNotReady    NotifyTimeResult = 3
)

type AccountStatus int

const (
	AccountStatusNormal  AccountStatus = 1
	AccountStatusDisable AccountStatus = 2
)

type CycleType int

const (
	CycleTypeDay   CycleType = 1
	CycleTypeWeek  CycleType = 2
	CycleTypeMonth CycleType = 3
)
