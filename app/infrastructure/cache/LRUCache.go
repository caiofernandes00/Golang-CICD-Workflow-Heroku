package cache

import (
	"overengineering-my-application/app/util"
	"time"
)

type CacheElement[T any] struct {
	node  *util.Node[string]
	value T
	ttl   time.Time
}

type LRUCache[T any] struct {
	dict       map[string]*CacheElement[T]
	linkedList *util.DoublyLinkedList[string]
	capacity   int
	ticker     *time.Ticker
}

func NewLRUCache[T any](capacity int) *LRUCache[T] {
	lru := &LRUCache[T]{
		dict:       make(map[string]*CacheElement[T]),
		linkedList: &util.DoublyLinkedList[string]{},
		capacity:   capacity + 1,
		ticker:     time.NewTicker(5 * time.Second),
	}

	go func() {
		for range lru.ticker.C {
			lru.checkTTL()
		}
	}()

	return lru
}

func (c *LRUCache[T]) Get(key string) (T, bool) {
	cacheValue, ok := c.dict[key]
	if !ok {
		var emptyValue T
		return emptyValue, false
	}

	c.linkedList.MoveToFront(cacheValue.node)
	return cacheValue.value, true
}

func (c *LRUCache[T]) Set(key string, value T, ttl time.Duration) {
	if cacheValue, ok := c.dict[key]; !ok {
		c.dict[key] = &CacheElement[T]{
			node:  c.linkedList.AddToFront(key),
			value: value,
			ttl:   time.Now().Add(ttl),
		}

		if len(c.dict) == c.capacity {
			lruValue := c.linkedList.Back()
			c.linkedList.RemoveTail()
			delete(c.dict, lruValue.Value)
		}
	} else {
		cacheValue.value = value
		c.linkedList.MoveToFront(cacheValue.node)
	}
}

func (c *LRUCache[T]) Evict(key string) {
	if cacheValue, ok := c.dict[key]; ok {
		c.linkedList.RemoveNode(cacheValue.node)
		delete(c.dict, key)
	}
}

func (c *LRUCache[T]) checkTTL() {
	for key, cacheValue := range c.dict {
		// check if ttl is zero
		if cacheValue.ttl.Before(time.Now()) && !cacheValue.ttl.IsZero() {
			c.Evict(key)
		}
	}
}
