package circleQueue

import (
	"errors"
	"fmt"
)

const QueueSize = 4

type CircleQueue struct {
	dataStore [QueueSize + 1]interface{}
	front     int
	rear      int
	len       int
}

func (q *CircleQueue) Front() int {
	return q.front
}

func (q *CircleQueue) Rear() int {
	return q.rear
}

func InitQueue(q *CircleQueue) {
	q.front, q.rear = 0, 0
	q.len = QueueSize + 1 // 循环队列最多只能存储QueueSize - 1 个数据; 最后一个数据表示满格
}

func (q *CircleQueue) EnQueue(data interface{}) error {
	// check if queue is full
	if (q.rear+1)%q.len == q.front%q.len {
		return errors.New("Full Queue")
	}

	q.dataStore[q.rear] = data
	q.rear = (q.rear + 1) % q.len
	return nil
}

func (q *CircleQueue) DeQueue() (data interface{}, err error) {
	// check queue is empty
	if q.front == q.rear {
		return nil, errors.New("Empty Queue")
	}

	res := q.dataStore[q.front]
	q.dataStore[q.front] = nil
	q.front = (q.front + 1) % q.len
	return res, nil
}

func (q *CircleQueue) Length() int {
	return (q.rear - q.front + q.len) % q.len
}

func (q *CircleQueue) Show() string {
	return fmt.Sprint(q.dataStore)
}
