package resilience

import (
	"errors"
	"sync"
	"time"
)

type State string

const (
	Closed   State = "closed"
	Open     State = "open"
	HalfOpen State = "half open"
)

type CircuitBreaker struct {
	failures  int
	interval  time.Duration
	threshold int
	state     State

	mu          sync.Mutex
	isExecuting bool
}

func NewCircuitBreaker(interval time.Duration, threshold int) *CircuitBreaker {
	return &CircuitBreaker{
		interval:    interval,
		threshold:   threshold,
		state:       Closed,
		mu:          sync.Mutex{},
		isExecuting: false,
	}
}

func (cb *CircuitBreaker) handleError() {
	cb.failures++
	if cb.failures > cb.threshold {
		cb.state = Open
		cb.Heal()
	}
}

func (cb *CircuitBreaker) Call(fn func() error) error {
	if cb.state == Open {
		return errors.New("circuit breaker is open")
	}

	err := fn()
	if err != nil {
		cb.handleError()
	}

	return err
}

func (cb *CircuitBreaker) State() State {
	return cb.state
}

func (cb *CircuitBreaker) Failures() int {
	return cb.failures
}

func (cb *CircuitBreaker) Heal() {
	if !cb.IsHealing() {
		cb.isExecuting = true
		go cb.decreaseFailures()
	}
}

func (cb *CircuitBreaker) IsHealing() bool {
	return cb.isExecuting
}

func (cb *CircuitBreaker) decreaseFailures() {
	ticker := time.NewTicker(cb.interval)
	finish := func() {
		ticker.Stop()
		cb.isExecuting = false
	}

	for range ticker.C {
		if cb.failures > 0 {
			cb.failures--
		}

		if cb.failures == 0 {
			cb.state = Closed
			finish()
			break
		}

		if cb.failures < cb.threshold {
			cb.state = HalfOpen
		}
	}
}
