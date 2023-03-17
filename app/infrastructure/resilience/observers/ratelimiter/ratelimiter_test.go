package ratelimiter

import (
	"overengineering-my-application/app/infrastructure/resilience"
	"overengineering-my-application/app/infrastructure/resilience/observable/circuitbreaker"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// Core functionality
func Test_CallFunctionWithSuccess(t *testing.T) {
	// Arrange
	rl := NewRateLimiter(1 * time.Second, 1 * time.Second)
	// Act
	err := callMultipleTimes(rl, 1)

	// Assert
	require.NoError(t, err)
	require.Equal(t, time.Now().Round(time.Second), rl.lastRequest.Round(time.Second))
}

func Test_CallFunctionWithFailure(t *testing.T) {
	// Arrange
	rl := NewRateLimiter(1 * time.Second, 1 * time.Second)
	// Act
	err := callMultipleTimes(rl, 2)
	// Assert
	require.Error(t, err)
	require.Equal(t, resilience.ErrRateLimitExceeded, err)
}

func callMultipleTimes(rl *RateLimiter, n int) error {
	var err error
	for i := 0; i < n; i++ {
		err = rl.Call(func() error {
			return nil
		})
	}

	return err
}

// Observer functionality
func Test_NotifyStateHalfOpen(t *testing.T) {
	// Arrange
	rl := NewRateLimiter(1 * time.Second, 1 * time.Second)
	changeState := circuitbreaker.ChangeState{From: circuitbreaker.Open, To: circuitbreaker.HalfOpen}
	// Act
	rl.Notify(changeState)
	// Assert
	require.Equal(t, rl.baseRateLimit, rl.dynamicRateLimit)
}
