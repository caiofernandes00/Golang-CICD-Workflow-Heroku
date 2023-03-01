package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SettingValue(t *testing.T) {
	cache := NewLRUCache[int](5, 100)

	cache.Set("key1", 1)
	cache.Set("key2", 2)
	cache.Set("key3", 3)
	cache.Set("key4", 4)

	value, ok := cache.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, 1, value)

	value, ok = cache.Get("key2")
	assert.True(t, ok)
	assert.Equal(t, 2, value)

	value, ok = cache.Get("key3")
	assert.True(t, ok)
	assert.Equal(t, 3, value)

	value, ok = cache.Get("key4")
	assert.True(t, ok)
	assert.Equal(t, 4, value)
}

func Test_SettingValueAboveThreshold(t *testing.T) {
	cache := NewLRUCache[int](5, 100)

	cache.Set("key1", 1)
	cache.Set("key2", 2)
	cache.Set("key3", 3)
	cache.Set("key4", 4)

	value, ok := cache.Get("key1")
	assert.False(t, ok)
	assert.Equal(t, 0, value)

	value, ok = cache.Get("key2")
	assert.True(t, ok)
	assert.Equal(t, 2, value)

	value, ok = cache.Get("key3")
	assert.True(t, ok)
	assert.Equal(t, 3, value)

	value, ok = cache.Get("key4")
	assert.True(t, ok)
	assert.Equal(t, 4, value)
}
