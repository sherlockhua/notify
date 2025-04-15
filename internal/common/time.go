package common

import "time"

const (
	TimeLayoutWithZone    = "2006-01-02 15:04:05 -0700 MST"
	TimeLayoutWithoutZone = "2006-01-02 15:04:05"
)

type TimeUtil interface {
	GetCurrentTime() time.Time
}

type TimeUtilImpl struct {
}

func NewTimeUtil() TimeUtil {
	return &TimeUtilImpl{}
}

func DaysBetween(start, end time.Time) int {
	start = start.UTC()
	end = end.UTC()

	start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, time.UTC)
	end = time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, time.UTC)

	return int(end.Sub(start).Hours() / 24)
}

func (*TimeUtilImpl) GetCurrentTime() time.Time {
	return time.Now()
}
