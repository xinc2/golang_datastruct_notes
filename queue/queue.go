package queue

import (
	"fmt"
)

type IQueue interface {
	Size() int
	Header() interface{}
	Tail() interface{}
	IsEmpty() bool
	IsFull() bool
	EnQueue(val interface{})
	DeQueue() interface{}
	Clear()
}

type Queue struct {
	dataStore []interface{}
	theSize   int
}

func NewQueue() *Queue {
	q := new(Queue)
	q.dataStore = make([]interface{}, 0)
	q.theSize = 0
	return q
}

func (q *Queue) Size() int {
	return q.theSize
}

func (q *Queue) Header() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.dataStore[0]
}

func (q *Queue) Tail() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.dataStore[q.theSize-1]
}

func (q *Queue) IsEmpty() bool {
	return q.theSize < 1
}

func (q *Queue) EnQueue(val interface{}) {
	q.dataStore = append(q.dataStore, val)
	q.theSize++
}

func (q *Queue) DeQueue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	val := q.dataStore[0]
	if q.theSize > 1 {
		q.dataStore = q.dataStore[1:]
	} else {
		q.dataStore = make([]interface{}, 0)
	}
	q.theSize--
	return val
}

func (q *Queue) Clear() {
	q.dataStore = make([]interface{}, 0)
	q.theSize = 0
}

func (q *Queue) Show() string {
	return fmt.Sprint(q.dataStore)
}
