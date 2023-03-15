package resilience

import (
	"overengineering-my-application/app/util"
)

type Observer interface {
	Notify(data interface{})
}

type Observable struct {
	subs *util.DoublyLinkedList[Observer]
}

func NewObservable() *Observable {
	return &Observable{
		subs: &util.DoublyLinkedList[Observer]{},
	}
}

func (o *Observable) Subscribe(s Observer) {
	o.subs.AddToFront(s)
}

func (o *Observable) Unsubscribe(s Observer) {
	o.subs.RemoveValue(s)
}

func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Back(); z != nil; z = z.Next() {
		z.Value.Notify(data)
	}
}
