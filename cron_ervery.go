package robin

import (
	"time"
)

var everyExecutor = NewEverySchedulerExecutor()

type everySchedulerExecutor struct {
	fiber Fiber
}

func NewEverySchedulerExecutor() *everySchedulerExecutor {
	return new(everySchedulerExecutor).init()
}

func (c *everySchedulerExecutor) init() *everySchedulerExecutor {
	c.fiber = NewGoroutineMulti()
	c.fiber.Start()
	return c
}

func EverySunday() *Job {
	return newCronWeekday(time.Sunday)
}

func EveryMonday() *Job {
	return newCronWeekday(time.Monday)
}

func EveryTuesday() *Job {
	return newCronWeekday(time.Tuesday)
}

func EveryWednesday() *Job {
	return newCronWeekday(time.Wednesday)
}

func EveryThursday() *Job {
	return newCronWeekday(time.Thursday)
}

func EveryFriday() *Job {
	return newCronWeekday(time.Friday)
}

func EverySaturday() *Job {
	return newCronWeekday(time.Saturday)
}

func newCronWeekday(weekday time.Weekday) *Job {
	c := NewJob(1, everyExecutor.fiber, delayNone)
	c.unit = weeks
	c.weekday = weekday
	return c
}