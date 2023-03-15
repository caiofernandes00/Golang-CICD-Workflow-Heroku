package circuitbreaker

import (
	"errors"
	resilience "overengineering-my-application/app/infrastructure/resilience/observable"
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
	resilience.Observable

	failures  int
	interval  time.Duration
	threshold int
	state     State

	mu          sync.Mutex
	isExecuting bool
}

func NewCircuitBreaker(interval time.Duration, threshold int) *CircuitBreaker {
	cb := &CircuitBreaker{
		interval:    interval,
		threshold:   threshold,
		state:       Closed,
		mu:          sync.Mutex{},
		isExecuting: false,
	}

	cb.Observable = *resilience.NewObservable()

	return cb
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

func (cb *CircuitBreaker) IsHealing() bool {
	return cb.isExecuting
}

func (cb *CircuitBreaker) heal() {
	if !cb.IsHealing() {
		cb.isExecuting = true
		go cb.decreaseFailures()
	}
}

func (cb *CircuitBreaker) handleError() {
	cb.failures++
	if cb.failures > cb.threshold {
		cb.changeStatus(Open)
		cb.heal()
	}
}

func (cb *CircuitBreaker) changeStatus(status State) {
	cb.state = status
	cb.Fire(ChangeState{From: cb.state, To: status})
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
			cb.changeStatus(Closed)
			finish()
			break
		}

		if cb.failures < cb.threshold {
			cb.changeStatus(HalfOpen)
		}
	}
}
