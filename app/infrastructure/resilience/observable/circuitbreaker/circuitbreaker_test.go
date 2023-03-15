package circuitbreaker

import (
	"errors"
	mock_resilience "overengineering-my-application/app/infrastructure/resilience/observable/circuitbreaker/mock"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

// Core functionality

func Test_CallAFuctionWithSuccess(t *testing.T) {
	cb := NewCircuitBreaker(0, 0)
	err := cb.Call(func() error {
		return nil
	})
	require.Nil(t, err)

	require.Equal(t, Closed, cb.State())
	require.Equal(t, 0, cb.failures)
}

func Test_CallAFuctionWithError(t *testing.T) {
	cb := NewCircuitBreaker(50, 1)
	err := execFuncCallWithErrorTimes(cb, 1, "generic error")

	require.Equal(t, "generic error", err.Error())
	require.Equal(t, Closed, cb.State())

	err = execFuncCallWithErrorTimes(cb, 2, "generic error")

	require.NotNil(t, err)
	require.Equal(t, "circuit breaker is open", err.Error())
	require.Equal(t, Open, cb.State())
	require.Equal(t, 2, cb.failures)
}

func Test_CallAFuctionWithErrorAndWaitForHalfState(t *testing.T) {
	cb := NewCircuitBreaker(2*time.Second, 3)
	err := execFuncCallWithErrorTimes(cb, 5, "generic error")

	require.Equal(t, Open, cb.State())
	require.Equal(t, 4, cb.failures)
	require.Equal(t, "circuit breaker is open", err.Error())

	time.Sleep(6 * time.Second)
	err = execFuncCallWithErrorTimes(cb, 1, "generic error")

	require.NotNil(t, err)
	require.Equal(t, "generic error", err.Error())
	require.Equal(t, HalfOpen, cb.State())
}

func execFuncCallWithErrorTimes(cb *CircuitBreaker, times int, msgError string) error {
	var err error
	for i := 0; i < times; i++ {
		err = cb.Call(func() error {
			return errors.New(msgError)
		})
	}
	return err
}

// Observable
func Test_Fire(t *testing.T) {
	cb := NewCircuitBreaker(0, 0)
	ctrl := gomock.NewController(t)
	mockObserver := mock_resilience.NewMockObserver(ctrl)
	mockObserver2 := mock_resilience.NewMockObserver(ctrl)
	mockObserver.EXPECT().Notify("test").Times(1)
	mockObserver2.EXPECT().Notify("test").Times(1)

	cb.Subscribe(mockObserver)
	cb.Subscribe(mockObserver2)

	cb.Fire("test")
}
