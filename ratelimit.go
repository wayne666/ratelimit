package ratelimit

import (
	"sync"
	"time"
	"unsafe"
)

type Limiter interface {
	Take() time.Time
}

type limiter struct {
	sync.Mutex
	last       time.Time
	sleepFor   time.Duration
	perRequest time.Duration
}

func New(rate int) Limiter {
	l := &limiter{
		perRequest: time.Second / time.Duration(rate),
	}

	return l
}

func (l *limiter) Take() time.Time {
	l.Lock()
	defer l.Unlock()

	now := time.Now()

	if unsafe.Sizeof(l.last) == 0 {
		l.last = now
		return l.last
	}

	l.sleepFor += l.perRequest - now.Sub(l.last)

	if l.sleepFor > 0 {
		time.Sleep(l.sleepFor)
		l.last = now.Add(l.sleepFor)
		l.sleepFor = 0
	} else {
		l.last = now
	}

	return l.last
}
