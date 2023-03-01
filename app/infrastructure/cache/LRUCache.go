package cache

import (
	"overengineering-my-application/app/util"
)

type CacheElement[T any] struct {
	node  *util.Node[string]
	value T
	ttl   int
}

type LRUCache[T any] struct {
	dict       map[string]*CacheElement[T]
	linkedList *util.DoublyLinkedList[string]
	capacity   int
}

func NewLRUCache[T any](capacity int) *LRUCache[T] {
	return &LRUCache[T]{
		dict:       make(map[string]*CacheElement[T]),
		linkedList: &util.DoublyLinkedList[string]{},
		capacity:   capacity,
	}
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

func (c *LRUCache[T]) Set(key string, value T) {
	if cacheValue, ok := c.dict[key]; !ok {
		c.dict[key] = &CacheElement[T]{
			node:  c.linkedList.AddToFront(key),
			value: value,
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
