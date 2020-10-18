package arraylist

import (
	"errors"
	"fmt"
)

// StackArrayX is a interface
type StackArrayX interface {
	Clear()
	Size() int
	Pop() (interface{}, error)
	Push(val interface{}) error
	IsFull() bool
	IsEmpty() bool
	String() string
}

// StackX is data struct
type StackX struct {
	arr       *ArrayList
	SIterator Iterator
}

// NewArrayListStackX is a construct func
func NewArrayListStackX() *StackX {
	s := new(StackX)
	s.arr = NewArrayList()
	s.SIterator = s.arr.Iterator()
	return s
}

// Clear is a method implemented to clean up the stack
func (s *StackX) Clear() {
	s.arr.Clear()
}

// Size is a method implemented to calculate the real length used in the stack
func (s *StackX) Size() int {
	return s.arr.theSize
}

// Pop is a method implemented to remove the last element
func (s *StackX) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Empty Stack, there's no available element in stack")
	}
	lastVal := s.arr.dataStore[s.arr.theSize-1]
	s.arr.Delete(s.arr.theSize - 1)
	return lastVal, nil

}

// Push is a method implemented to append an element to the stack
func (s *StackX) Push(val interface{}) error {
	if s.IsFull() {
		return errors.New("Stack is full, there's no extra available spaces")
	}
	s.arr.Append(val)
	return nil
}

// IsFull is a method implemented to make sure if the stack is full
func (s *StackX) IsFull() bool {
	return s.arr.theSize >= 10
}

//IsEmpty is a method implemented to make sure if there's no any element in stack
func (s *StackX) IsEmpty() bool {
	return s.arr.theSize == 0
}

// String is a method implemented to show stack as type string
func (s *StackX) String() string {
	return fmt.Sprint(s.arr.dataStore)
}
