package circuitbreaker

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
	failures  int
	interval  time.Duration
	threshold int
	state     State
}

func NewCircuitBreaker(interval time.Duration, threshold int) *CircuitBreaker {
	return &CircuitBreaker{
		interval:  interval,
		threshold: threshold,
		state:     Closed,
	}
}

func (cb *CircuitBreaker) HandleError() {
	cb.failures++
	if cb.failures > cb.threshold {
		cb.state = Open
		go cb.decreaseFailures()
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

func (cb *CircuitBreaker) decreaseFailures() {
	ticker := time.NewTicker(cb.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			cb.failures--

			if cb.failures < cb.threshold {
				cb.state = HalfOpen
			}
		}
	}
}
