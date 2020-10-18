package linkList

type ILinkStack interface {
	IsEmpty() bool
	Length() int
	Push(data interface{})
	Pop() interface{}
}

type Node struct {
	data  interface{}
	pNext *Node
}

func NewLinkStack() *Node {
	return &Node{}
}

func (n *Node) IsEmpty() bool {
	if n.pNext == nil {
		return true
	}
	return false
}

func (n *Node) Length() int {
	lenght := 0
	for n.pNext != nil {
		n = n.pNext
		lenght++
	}
	return lenght
}

// 链表栈首部插入元素
func (n *Node) Push(val interface{}) {
	newNode := new(Node)
	newNode.data = val
	newNode.pNext = n.pNext
	n.pNext = newNode
}

// 链表栈的首部移除元素
func (n *Node) Pop() interface{} {
	if n.IsEmpty() {
		return nil
	}

	val := n.pNext.data
	n.pNext = n.pNext.pNext

	return val
}
