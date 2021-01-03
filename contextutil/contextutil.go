package contextutil

import (
	"context"
	"time"
)

// ctx, cancel, err := context.WithDeadline(tkinfo.WaitingSeatTimeoutDuration)
func WithDeadline(durationString string) (context.Context, context.CancelFunc, error) {
	var duration time.Duration
	duration, err := time.ParseDuration(durationString)
	if err != nil {
		return nil, nil, err
	}
	d := time.Now().Add(duration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	return ctx, cancel, nil
}
