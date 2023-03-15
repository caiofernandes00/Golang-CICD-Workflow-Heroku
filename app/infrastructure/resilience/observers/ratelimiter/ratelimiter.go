package ratelimiter

import (
	"log"
	"math"
	"overengineering-my-application/app/infrastructure/resilience"
	"overengineering-my-application/app/infrastructure/resilience/observable/circuitbreaker"
	"time"
)

type RateLimiter struct {
	rateLimit        time.Duration
	dynamicRateLimit time.Duration
	lastRequest      time.Time
}

func NewRateLimiter(rateLimit time.Duration) *RateLimiter {
	return &RateLimiter{
		rateLimit:        rateLimit,
		dynamicRateLimit: rateLimit,
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
		case circuitbreaker.Open:
			rl.dynamicRateLimit = rl.rateLimit * 2
		case circuitbreaker.HalfOpen:
			rl.dynamicRateLimit = rl.rateLimit * time.Duration(math.Floor(float64(rl.rateLimit)*1.5))
		default:
			rl.dynamicRateLimit = rl.rateLimit
		}
	} else {
		log.Printf("unexpected data type: %T", data)
	}
}
