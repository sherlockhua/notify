package task_strategy

import (
	"context"
	"fmt"
	"notify/internal/common"
)

const (
	TaskTriggerDiffSeconds = 30
)

type TaskStrategy interface {
	IsTimeToNotify(ctx context.Context) common.NotifyTimeResult
}

type TaskStrategyFactory interface {
	CreateTaskStrategy(ctx context.Context, taskType common.TaskType, strategyData string) (TaskStrategy, error)
}

type taskStrategyFactory struct {
	timeUtil common.TimeUtil
}

func NewTaskStrategyFactory(timeUtil common.TimeUtil) TaskStrategyFactory {
	return &taskStrategyFactory{
		timeUtil: timeUtil,
	}
}

func (f *taskStrategyFactory) CreateTaskStrategy(ctx context.Context, taskType common.TaskType, strategyData string) (TaskStrategy, error) {
	switch taskType {
	case common.TaskTypeOnce:
		return NewOnceStrategy(strategyData, f.timeUtil)
	case common.TaskTypeDayCycle:
		return NewDayCycleStrategy(strategyData, f.timeUtil)
	case common.TaskTypeWeekCycle:
		//return NewWeekCycleStrategy(task.StrategyData), nil
	case common.TaskTypeMonthCycle:
		//return NewMonthCycleStrategy(task.StrategyData), nil
	case common.TaskTypeYearCycle:
		//return NewYearCycleStrategy(task.StrategyData), nil
	default:
		return nil, common.ErrInvalidTaskType
	}
	return nil, fmt.Errorf("unsupported task type: %v", taskType)
}
