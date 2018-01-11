package ratelimit

import (
	"sync"
	"time"
)

type limiter struct {
	sync.Mutex
	last       time.Duration
	sleepFor   time.Duration
	perRequest time.Duration
}

func New(rate int) *limiter {
	l := &limiter{
		perRequest: time.Second / time.Duration(rate),
	}

	return l
}

func (l *limiter) Take() Time.time {
	return
}
