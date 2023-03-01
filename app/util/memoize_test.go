package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_memoization(t *testing.T) {
	fakeFn := func(num1 int, num2 int) int {
		time.Sleep(1 * time.Second)
		return num1 + num2
	}

	memoizedFn := Memoize(fakeFn)

	start := time.Now()
	result := memoizedFn(1, 2)
	require.Equal(t, 3, result)
	require.True(t, time.Since(start) > 1*time.Second)

	start = time.Now()
	result = memoizedFn(1, 2)
	require.Equal(t, 3, result)
	require.True(t, time.Since(start) < 1*time.Second)
}
