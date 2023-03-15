package resilience

import (
	"time"
)

type RateLimiter struct {
	rateLimit   time.Duration
	lastRequest time.Time
}

func NewRateLimiter(rateLimit time.Duration) *RateLimiter {
	return &RateLimiter{
		rateLimit: rateLimit,
	}
}

func (rl *RateLimiter) Call(fn func() error) error {
	if rl.lastRequest.IsZero() {
		rl.lastRequest = time.Now()
		return fn()
	}
	
	if time.Since(rl.lastRequest)*time.Second < rl.rateLimit {
		return ErrRateLimitExceeded
	}

	rl.lastRequest = time.Now()

	return fn()
}
