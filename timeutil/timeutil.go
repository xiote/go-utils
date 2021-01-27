package timeutil

import (
	"time"
)

type Clock struct {
	SetTime   time.Time
	StartTime time.Time
	Slowness  time.Duration
}

func NewClock(settime time.Time) *Clock {
	now := time.Now()
	return &Clock{settime, now, now.Sub(settime)}
}

func (c *Clock) AlarmAt(alarmtime time.Time) <-chan time.Time {
	duration := alarmtime.Sub(c.Now())

	return time.After(duration)
}

func (c *Clock) Now() time.Time {
	return c.SetTime.Add(time.Since(c.StartTime))
}

func MustParseDuration(s string) time.Duration {
	if d, err := time.ParseDuration(s); err != nil {
		panic(err)
	} else {
		return d
	}

}
