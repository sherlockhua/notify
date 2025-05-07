package common

type TaskStatus int

const (
	TaskStatusRunning  TaskStatus = 1
	TaskStatusFinished TaskStatus = 2
	TaskStatusCannel   TaskStatus = 3
	TaskStatusPaused   TaskStatus = 4
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
	NotifyTimeResultUnknown         NotifyTimeResult = 0
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

type TemplateType int

const (
	TemplateTypeHoliday  TemplateType = 1
	TemplateTypeMedicine TemplateType = 2
)

type TemplateStatus int

const (
	TemplateStatusNormal TemplateStatus = 1
	TemplateStatusDelete TemplateStatus = 2
)
