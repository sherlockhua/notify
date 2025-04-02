package common

type TaskStatus int

const (
	TaskStatusNotStart TaskStatus = 1
	TaskStatusRunning  TaskStatus = 2
	TaskStatusFinished TaskStatus = 3
	TaskStatusCannel   TaskStatus = 4
)

type AccountStatus int

const (
	AccountStatusNormal  AccountStatus = 1
	AccountStatusDisable AccountStatus = 2
)
