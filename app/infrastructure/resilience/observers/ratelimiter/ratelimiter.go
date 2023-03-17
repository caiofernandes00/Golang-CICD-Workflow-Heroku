package ratelimiter

import (
	"log"
	"math"
	"overengineering-my-application/app/infrastructure/resilience"
	"overengineering-my-application/app/infrastructure/resilience/observable/circuitbreaker"
	"time"
)

type RateLimiter struct {
	baseRateLimit            time.Duration
	dynamicRateLimit         time.Duration
	lastRequest              time.Time
	baseExponentialFactor    time.Duration
	dynamicExponentialFactor time.Duration
}

func NewRateLimiter(baseRateLimit time.Duration, baseExponentialFactor time.Duration) *RateLimiter {
	return &RateLimiter{
		baseRateLimit:            baseRateLimit,
		dynamicRateLimit:         baseRateLimit,
		baseExponentialFactor:    baseExponentialFactor,
		dynamicExponentialFactor: baseExponentialFactor,
	}
}

func (rl *RateLimiter) Call(fn func() error) error {
	if rl.lastRequest.IsZero() {
		rl.lastRequest = time.Now()
		return fn()
	}

	if time.Since(rl.lastRequest)*time.Second < rl.dynamicRateLimit {
		return resilience.ErrRateLimitExceeded
	}

	rl.lastRequest = time.Now()

	return fn()
}

func (rl *RateLimiter) Notify(data interface{}) {
	if val, ok := data.(circuitbreaker.ChangeState); ok {
		switch val.To {
		case circuitbreaker.HalfOpen:
			rl.dynamicRateLimit = rl.exponentialBackoff()
		default:
			rl.dynamicExponentialFactor = rl.baseExponentialFactor
			rl.dynamicRateLimit = rl.baseRateLimit
		}
	} else {
		log.Printf("unexpected data type: %T", data)
	}
}

// TODO: Find a better way to calculate the dynamic rate limit
func (rl *RateLimiter) exponentialBackoff() time.Duration {
	rl.dynamicExponentialFactor++
	return rl.dynamicRateLimit * time.Duration(math.Pow(float64(rl.baseExponentialFactor), 2))
}
