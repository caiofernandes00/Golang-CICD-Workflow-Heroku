package resilience

import (
	"errors"
	"time"
)

type State string

const (
	Closed   State = "closed"
	Open     State = "open"
	HalfOpen State = "half open"
)

type CircuitBreaker struct {
	failures                    int
	interval                    time.Duration
	threshold                   int
	state                       State
	runningDecreaseFailuresChan chan bool
}

func NewCircuitBreaker(interval time.Duration, threshold int) *CircuitBreaker {
	return &CircuitBreaker{
		interval:                    interval,
		threshold:                   threshold,
		state:                       Closed,
		runningDecreaseFailuresChan: make(chan bool),
	}
}

func (cb *CircuitBreaker) HandleError() {
	cb.failures++
	if cb.failures > cb.threshold {
		cb.state = Open
		if !<-cb.runningDecreaseFailuresChan {
			go cb.decreaseFailures()
		}
	}
}

func (cb *CircuitBreaker) Call(fn func() error) error {
	if cb.state == Open {
		return errors.New("circuit breaker is open")
	}

	err := fn()
	if err != nil {
		cb.HandleError()
	} else if cb.state == HalfOpen {
		cb.state = Closed
	}

	return err
}

func (cb *CircuitBreaker) State() State {
	return cb.state
}

func (cb *CircuitBreaker) decreaseFailures() {
	ticker := time.NewTicker(cb.interval)
	defer ticker.Stop()

	for range ticker.C {
		if cb.failures > 0 {
			cb.failures--
		}

		if cb.failures == 0 {
			cb.state = Closed
			cb.runningDecreaseFailuresChan <- false
			break
		}

		if cb.failures < cb.threshold {
			cb.state = HalfOpen
		}
	}
}
