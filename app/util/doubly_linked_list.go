package util

type Node[T comparable] struct {
	Value T
	prev  *Node[T]
	next  *Node[T]
}

type DoublyLinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	len  int
}

func (l *DoublyLinkedList[T]) AddToFront(value T) *Node[T] {
	newNode := &Node[T]{Value: value}

	if l.tail == nil {
		l.tail = newNode
		l.head = newNode
	} else {
		l.head.next = newNode
		newNode.prev = l.head
		l.head = newNode
	}

	l.len++

	return newNode
}

func (l *DoublyLinkedList[T]) RemoveValue(value T) {
	if l.IsEmpty() {
		return
	}

	node := l.head
	for node != nil {
		if node.Value == value {
			l.RemoveNode(node)
			return
		}
		node = node.prev
	}
}

func (l *DoublyLinkedList[T]) RemoveNode(node *Node[T]) {
	if node == nil {
		return
	}

	if node == l.head {
		l.head = node.prev
	} else {
		node.next.prev = node.prev
	}

	if node == l.tail {
		l.tail = node.next
	} else {
		node.prev.next = node.next
	}

	node.prev = nil
	node.next = nil

	l.len--
}

func (l *DoublyLinkedList[T]) MoveToFront(node *Node[T]) {
	if node == nil {
		return
	}

	if node == l.head {
		return
	}

	if node == l.tail {
		l.tail = node.next
		l.tail.prev = nil
	} else {
		node.prev.next = node.next
	}

	node.next.prev = node.prev

	node.prev = l.head
	node.next = nil
	l.head.next = node
	l.head = node
}

func (l *DoublyLinkedList[T]) MoveToBack(node *Node[T]) {
	if node == nil {
		return
	}

	if node == l.tail {
		return
	}

	if node == l.head {
		l.head = node.prev
		l.head.next = nil
	} else {
		node.next.prev = node.prev
	}

	node.prev.next = node.next

	node.prev = nil
	node.next = l.tail
	l.tail.prev = node
	l.tail = node
}

func (l *DoublyLinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l *DoublyLinkedList[T]) Front() *Node[T] {
	return l.head
}

func (l *DoublyLinkedList[T]) RemoveHead() {
	if l.IsEmpty() {
		return
	}

	l.RemoveNode(l.head)
}

func (l *DoublyLinkedList[T]) RemoveTail() {
	if l.IsEmpty() {
		return
	}

	l.RemoveNode(l.tail)
}

func (l *DoublyLinkedList[T]) Contains(value T) bool {
	if l.IsEmpty() {
		return false
	}

	node := l.head
	for node != nil {
		if node.Value == value {
			return true
		}
		node = node.prev
	}
	return false
}

func (l *DoublyLinkedList[T]) IsEmpty() bool {
	if l.tail == nil {
		return true
	}
	return false
}

func (l *DoublyLinkedList[T]) Len() int {
	return l.len
}

func (l *DoublyLinkedList[T]) Iterate() {
	if l.IsEmpty() {
		return
	}

	node := l.head
	for node != nil {
		node = node.prev
	}
}

func (l *DoublyLinkedList[T]) IterateReverse() {
	if l.IsEmpty() {
		return
	}

	node := l.tail
	for node != nil {
		node = node.next
	}
}
