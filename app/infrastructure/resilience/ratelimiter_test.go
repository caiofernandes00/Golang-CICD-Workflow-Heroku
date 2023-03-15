package resilience

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_CallFunctionWithSuccess(t *testing.T) {
	// Arrange
	rl := NewRateLimiter(1 * time.Second)
	// Act
	err := rl.Call(func() error {
		return nil
	})
	// Assert
	require.NoError(t, err)
	require.Equal(t, time.Now().Round(time.Second), rl.lastRequest.Round(time.Second))
}

func Test_CallFunctionWithFailure(t *testing.T) {
	// Arrange
	rl := NewRateLimiter(1)
	// Act
	err := callMultipleTimes(rl, 2)
	// Assert
	require.Error(t, err)
	require.Equal(t, ErrRateLimitExceeded, err)
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
