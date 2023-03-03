package resilience

import (
	"time"
)

type RateLimiter struct {
	rateLimit      time.Duration
	lastRequest    time.Time
	circuitBreaker *CircuitBreaker
}

func NewRateLimiter(rateLimit time.Duration, circuitBreaker *CircuitBreaker) *RateLimiter {
	return &RateLimiter{
		rateLimit:      rateLimit,
		circuitBreaker: circuitBreaker,
	}
}

func (rl *RateLimiter) Call(fn func() error) error {
	if time.Since(rl.lastRequest)*time.Second < rl.rateLimit*time.Second {
		return ErrRateLimitExceeded
	}
	if rl.circuitBreaker.State() == HalfOpen {
		rl.rateLimit = rl.rateLimit * 2
	}

	rl.lastRequest = time.Now()

	return fn()
}
