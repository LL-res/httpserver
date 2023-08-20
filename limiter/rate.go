package limiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	sync.Mutex
	rate        int
	cap         int
	cur         int
	lastRequest time.Time
}

func NewR(rate int, cap int) *RateLimiter {
	return &RateLimiter{
		rate: rate,
		cap:  cap,
	}
}

func (l *RateLimiter) Allow() bool {
	l.Lock()
	defer l.Unlock()
	now := time.Now()
	l.cur += int(now.Sub(l.lastRequest).Seconds()) * l.rate
	l.lastRequest = now
	if l.cur > l.cap {
		l.cur = l.cap
	}
	if l.cur > 0 {
		l.cur--
		return true
	}
	return false
}
