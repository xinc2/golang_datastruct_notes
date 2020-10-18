package main

import (
	"data_struct/arraylist"
	"data_struct/circleQueue"
	"data_struct/linkList"
	"data_struct/queue"
	"data_struct/stackarray"
	"fmt"
)

// test for arraylist
func main1() {
	list := arraylist.NewArrayList()
	list.Append(1)
	list.Append("hello-world")
	list.Append(true)
	list.Append(1.001)
	fmt.Println(list.Size())
	fmt.Println(list.Get(2))
	fmt.Println(list.String())
	for i := 0; i < 10; i++ {
		list.Insert(1, "321")
		fmt.Println(list)
	}

}

// test arraylist iterator
func main2() {
	list := arraylist.NewArrayList()
	list.Append(1)
	list.Append("hello-world")
	list.Append(true)
	list.Append(1.001)
	list.Append("1001")
	fmt.Println(list)
	for it := list.Iterator(); it.HasNext(); {
		item, _ := it.Next()
		if item == "hello-world" {
			it.Remove()
		}
	}
	fmt.Println(list)
}

// test Stack
func main3() {
	s := stackarray.NewStack()
	s.Push(1)
	s.Push("hello-world")
	s.Push(true)
	s.Push(13.01)
	fmt.Println(s.String())
	fmt.Println(s.Pop())
	fmt.Println(s.String())
	for i := s.Size(); i >= 0; i-- {
		if _, err := s.Pop(); err != nil {
			break
		}
		fmt.Println(s.IsEmpty())
		fmt.Println(s.String())
	}
	for i := 0; i <= s.MaxSize(); i++ {
		if err := s.Push(1); err != nil {
			break
		}
		fmt.Println(s.IsFull())
		fmt.Println(s.String())
	}

}

// test Stack
func main4() {
	s := arraylist.NewArrayListStack()
	s.Push(1)
	s.Push("hello-world")
	s.Push(true)
	s.Push(13.01)
	fmt.Println(s.String())
	fmt.Println(s.Pop())
	fmt.Println(s.String())
	for i := s.Size(); i >= 0; i-- {
		if _, err := s.Pop(); err != nil {
			break
		}
		fmt.Println(s.IsEmpty())
		fmt.Println(s.String())
	}
	for i := 0; i <= s.MaxSize(); i++ {
		if err := s.Push(1); err != nil {
			break
		}
		fmt.Println(s.IsFull())
		fmt.Println(s.String())
	}

}

// test Stack iterator
func main5() {
	s := arraylist.NewArrayListStackX()
	s.Push(1)
	s.Push("hello-world")
	s.Push(true)
	s.Push(13.01)
	fmt.Println(s.String())
	for it := s.SIterator; it.HasNext(); {
		item, _ := it.Next()
		if item == "hello-world" {
			it.Remove()
		}
	}
	fmt.Println(s.String())
}

func main6() {
	q := queue.NewQueue()
	q.EnQueue(1)
	q.EnQueue("hello-world")
	q.EnQueue(true)
	q.EnQueue(13.01)
	fmt.Println("Queue before->", q.Show(), q.Size())
	// fmt.Println(q.DeQueue())
	// fmt.Println(q.DeQueue())
	// fmt.Println(q.DeQueue())
	// fmt.Println(q.DeQueue())
	length := q.Size()
	for i := 0; i < length; i++ {
		fmt.Println(q.DeQueue())
	}
	fmt.Println("Queue after-->", q.Show())
}

func main7() {
	var q circleQueue.CircleQueue
	circleQueue.InitQueue(&q)
	q.EnQueue(1)
	q.EnQueue("hello-world")
	q.EnQueue(true)
	q.EnQueue(13.53)
	q.EnQueue(14)
	fmt.Println("Queue before->", q.Show(), q.Length(), q.Front(), q.Rear())
	q.DeQueue()
	q.DeQueue()
	fmt.Println("Queue after->", q.Show(), q.Length(), q.Front(), q.Rear())
}

func main8() {
	linkstack := linkList.NewLinkStack()
	for i := 0; i < 10; i++ {
		linkstack.Push(i)
	}
	fmt.Println("Pop element....")
	for !linkstack.IsEmpty() {
		data := linkstack.Pop()
		fmt.Println(data)
	}
}

func main() {
	linkq := linkList.NewLinkQueue()
	for i := 0; i < 10; i++ {
		linkq.EnQueue(i)

	}
	// fmt.Println(linkq.Length())
	fmt.Println("Pop element....")
	for !linkq.IsEmpty() {
		data := linkq.DeQueue()
		fmt.Println(data)
	}
}
