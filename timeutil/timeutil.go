package timeutil

import (
	"time"
)

func MustParseDuration(s string) time.Duration {
	if d, err := time.ParseDuration(s); err != nil {
		panic(err)
	} else {
		return d
	}

}
