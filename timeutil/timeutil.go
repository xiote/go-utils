package timeutil

import (
	"strings"
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

func MustParseTime(s string) time.Time {
	s = strings.ReplaceAll(s, " ", "")
	if strings.HasSuffix(s, "시") {
		s = s + "00분"
	}

	if strings.HasPrefix(s, "오전") || strings.HasPrefix(s, "오후") {
		s = strings.ReplaceAll(s, "오전", "AM")
		s = strings.ReplaceAll(s, "오후", "PM")
		if t, err := time.Parse("PM3시04분", s); err != nil {
			panic(err)
		} else {
			return t
		}
	} else {
		if t, err := time.Parse("15시04분", s); err != nil {
			panic(err)
		} else {
			return t
		}
	}
}
