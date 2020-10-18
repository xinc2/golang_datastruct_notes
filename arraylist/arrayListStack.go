package arraylist

import (
	"errors"
	"fmt"
)

// StackArray is a interface
type StackArray interface {
	Clear()
	Size() int
	Pop() (interface{}, error)
	Push(val interface{}) error
	IsFull() bool
	IsEmpty() bool
	String() string
}

// Stack is data struct
type Stack struct {
	arr     *ArrayList
	capSize int // max capacity of the stack
}

// NewStack is a construct func
func NewArrayListStack() *Stack {
	s := new(Stack)
	s.arr = NewArrayList()
	s.capSize = 10
	return s
}

// Clear is a method implemented to clean up the stack
func (s *Stack) Clear() {
	s.arr.Clear()
	s.capSize = 0
}

// Size is a method implemented to calculate the real length used in the stack
func (s *Stack) Size() int {
	return s.arr.theSize
}

// MaxSize used to show the max capacity of stack
func (s *Stack) MaxSize() int {
	return s.capSize
}

// Pop is a method implemented to remove the last element
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Empty Stack, there's no available element in stack")
	}
	lastVal := s.arr.dataStore[s.arr.theSize-1]
	s.arr.Delete(s.arr.theSize - 1)
	return lastVal, nil

}

// Push is a method implemented to append an element to the stack
func (s *Stack) Push(val interface{}) error {
	if s.IsFull() {
		return errors.New("Stack is full, there's no extra available spaces")
	}
	s.arr.Append(val)
	return nil
}

// IsFull is a method implemented to make sure if the stack is full
func (s *Stack) IsFull() bool {
	return s.arr.theSize >= s.capSize
}

//IsEmpty is a method implemented to make sure if there's no any element in stack
func (s *Stack) IsEmpty() bool {
	return s.arr.theSize == 0
}

// String is a method implemented to show stack as type string
func (s *Stack) String() string {
	return fmt.Sprint(s.arr.dataStore)
}
