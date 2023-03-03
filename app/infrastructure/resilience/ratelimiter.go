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

func (rl *RateLimiter) Allow() (bool, error) {
	if time.Since(rl.lastRequest) < rl.rateLimit {
		return false, ErrRateLimitExceeded
	}

	rl.lastRequest = time.Now()

	return true, nil
}
