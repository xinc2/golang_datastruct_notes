package linkList

type ILinkQueue interface {
	IsEmpty() bool
	Length() int
	EnQueue(val interface{})
	DeQueue() interface{}
}

//链式队列，用来指明队列的首尾节点位置
type LinkQueue struct {
	front *Node
	rear  *Node
}

func NewLinkQueue() *LinkQueue {
	return &LinkQueue{}
}

func (lq *LinkQueue) IsEmpty() bool {
	if lq.front == nil {
		return true
	}
	return false
}

func (lq *LinkQueue) Length() int {
	length := 0
	for lq.front != nil {
		lq.front = lq.front.pNext
		length++
	}
	return length
}

func (lq *LinkQueue) EnQueue(val interface{}) {
	newNode := &Node{val, nil}
	if lq.front == nil {
		lq.front = newNode
		lq.rear = newNode
	} else {
		lq.rear.pNext = newNode
		lq.rear = lq.rear.pNext
	}
}

func (lq *LinkQueue) DeQueue() interface{} {
	if lq.front == nil {
		return nil
	}

	delNode := lq.front
	if lq.front == lq.rear { // 队列只有一个元素的情况
		lq.front = nil
		lq.rear = nil
	} else {
		lq.front = lq.front.pNext
	}

	return delNode.data
}
