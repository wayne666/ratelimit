package ratelimit

import (
	"sync"
	"time"
)

type Limiter interface {
	Take() Time.time
}

type limiter struct {
	sync.Mutex
	last       time.Duration
	sleepFor   time.Duration
	perRequest time.Duration
}

func New(rate int) Limiter {
	l := &limiter{
		perRequest: time.Second / time.Duration(rate),
	}

	return l
}

func (l *limiter) Take() Time.time {
	l.Lock()
	defer l.Unlock()

	now := time.Now()

	if l.last == nil {
		l.last = now
		return l.last
	}

	l.sleepFor += l.perRequest - now.Sub(l.last)

	if l.sleepFor > 0 {
		time.Sleep(l.sleepFor)
		l.last = now.Add(t.sleepFor)
		l.sleepFor = 0
	} else {
		l.last = now
	}

	return l.last
}
