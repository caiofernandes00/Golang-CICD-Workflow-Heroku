package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_SettingValue(t *testing.T) {
	cache := NewLRUCache[int](4)

	cache.Set("key1", 1, 0)
	cache.Set("key2", 2, 0)
	cache.Set("key3", 3, 0)
	cache.Set("key4", 4, 0)

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
	cache := NewLRUCache[int](3)

	cache.Set("key1", 1, 0)
	cache.Set("key2", 2, 0)
	cache.Set("key3", 3, 0)
	cache.Set("key4", 4, 0)

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

func Test_TTL(t *testing.T) {
	cache := NewLRUCache[int](4)

	cache.Set("key1", 1, 1)
	cache.Set("key2", 2, 1)
	cache.Set("key3", 3, 1)
	cache.Set("key4", 4, 1)

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

	time.Sleep(6 * time.Second)

	value, ok = cache.Get("key1")
	assert.False(t, ok)
	assert.Equal(t, 0, value)

	value, ok = cache.Get("key2")
	assert.False(t, ok)
	assert.Equal(t, 0, value)

	value, ok = cache.Get("key3")
	assert.False(t, ok)
	assert.Equal(t, 0, value)

	value, ok = cache.Get("key4")
	assert.False(t, ok)
	assert.Equal(t, 0, value)
}
