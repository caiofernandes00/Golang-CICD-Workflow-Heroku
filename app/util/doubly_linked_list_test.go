package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddToFront(t *testing.T) {
	l := DoublyLinkedList[int]{}
	l.AddToFront(1)
	l.AddToFront(2)
	l.AddToFront(3)
	l.AddToFront(4)

	assert.Equal(t, 4, l.len)

	assert.Equal(t, 4, l.head.Value)
	assert.Equal(t, 1, l.tail.Value)

	assert.Equal(t, 3, l.head.prev.Value)
	assert.Equal(t, 2, l.head.prev.prev.Value)
	assert.Equal(t, 1, l.head.prev.prev.prev.Value)

	assert.Equal(t, 2, l.tail.next.Value)
	assert.Equal(t, 3, l.tail.next.next.Value)
	assert.Equal(t, 4, l.tail.next.next.next.Value)
}

func Test_AddToBack(t *testing.T) {
	l := DoublyLinkedList[int]{}
	l.AddToBack(1)
	l.AddToBack(2)
	l.AddToBack(3)
	l.AddToBack(4)

	assert.Equal(t, 4, l.len)

	assert.Equal(t, 1, l.head.Value)
	assert.Equal(t, 4, l.tail.Value)

	assert.Equal(t, 2, l.head.prev.Value)
	assert.Equal(t, 3, l.head.prev.prev.Value)
	assert.Equal(t, 4, l.head.prev.prev.prev.Value)

	assert.Equal(t, 3, l.tail.next.Value)
	assert.Equal(t, 2, l.tail.next.next.Value)
	assert.Equal(t, 1, l.tail.next.next.next.Value)
}

func Test_MoveToFront(t *testing.T) {
	l := DoublyLinkedList[int]{}

	l.AddToFront(1)
	l.AddToFront(2)
	l.AddToFront(3)
	l.AddToFront(4)

	assert.Equal(t, 4, l.head.Value)
	assert.Equal(t, 3, l.head.prev.Value)
	assert.Equal(t, 2, l.head.prev.prev.Value)
	assert.Equal(t, 1, l.head.prev.prev.prev.Value)
	l.MoveToFront(l.head.prev.prev)
	assert.Equal(t, 2, l.head.Value)
	assert.Equal(t, 4, l.head.prev.Value)
	assert.Equal(t, 3, l.head.prev.prev.Value)
	assert.Equal(t, 1, l.head.prev.prev.prev.Value)
}

func Test_MoveToBack(t *testing.T) {
	l := DoublyLinkedList[int]{}

	l.AddToFront(1)
	l.AddToFront(2)
	l.AddToFront(3)
	l.AddToFront(4)

	assert.Equal(t, 4, l.head.Value)
	assert.Equal(t, 3, l.head.prev.Value)
	assert.Equal(t, 2, l.head.prev.prev.Value)
	assert.Equal(t, 1, l.head.prev.prev.prev.Value)
	l.MoveToBack(l.head.prev)
	assert.Equal(t, 4, l.head.Value)
	assert.Equal(t, 2, l.head.prev.Value)
	assert.Equal(t, 1, l.head.prev.prev.Value)
	assert.Equal(t, 3, l.head.prev.prev.prev.Value)
}

func Test_RemoveNode(t *testing.T) {
	l := DoublyLinkedList[int]{}

	l.AddToFront(1)
	l.AddToFront(2)
	l.AddToFront(3)
	l.AddToFront(4)

	assert.Equal(t, 4, l.head.Value)
	assert.Equal(t, 3, l.head.prev.Value)
	assert.Equal(t, 2, l.head.prev.prev.Value)
	assert.Equal(t, 1, l.head.prev.prev.prev.Value)
	l.RemoveNode(l.head.prev)
	assert.Equal(t, 4, l.head.Value)
	assert.Equal(t, 2, l.head.prev.Value)
	assert.Equal(t, 1, l.head.prev.prev.Value)
}

func Test_RemoveTail(t *testing.T) {
	l := DoublyLinkedList[int]{}

	l.AddToFront(1)
	l.AddToFront(2)
	l.AddToFront(3)
	l.AddToFront(4)

	assert.Equal(t, 1, l.tail.Value)
	assert.Equal(t, 2, l.tail.next.Value)
	assert.Equal(t, 3, l.tail.next.next.Value)
	assert.Equal(t, 4, l.tail.next.next.next.Value)
	l.RemoveTail()
	assert.Equal(t, 2, l.tail.Value)
	assert.Equal(t, 3, l.tail.next.Value)
	assert.Equal(t, 4, l.tail.next.next.Value)
}

func Test_RemoveHead(t *testing.T) {
	l := DoublyLinkedList[int]{}

	l.AddToFront(1)
	l.AddToFront(2)
	l.AddToFront(3)
	l.AddToFront(4)

	assert.Equal(t, 4, l.head.Value)
	assert.Equal(t, 3, l.head.prev.Value)
	assert.Equal(t, 2, l.head.prev.prev.Value)
	assert.Equal(t, 1, l.head.prev.prev.prev.Value)
	l.RemoveHead()
	assert.Equal(t, 3, l.head.Value)
	assert.Equal(t, 2, l.head.prev.Value)
	assert.Equal(t, 1, l.head.prev.prev.Value)
}
