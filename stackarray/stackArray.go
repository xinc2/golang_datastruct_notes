package stackarray

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
	dataStore []interface{}
	capSize   int // max capacity of the stack
	cursor    int // current length of the stack
}

// NewStack is a construct func
func NewStack() *Stack {
	s := new(Stack)
	s.dataStore = make([]interface{}, 0, 1000)
	s.capSize = 1000
	s.cursor = 0
	return s
}

// Clear is a method implemented to clean up the stack
func (s *Stack) Clear() {
	s.dataStore = make([]interface{}, 0, 10000)
	s.capSize = 10000
	s.cursor = 0
}

// Size is a method implemented to calculate the real length used in the stack
func (s *Stack) Size() int {
	return s.cursor
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
	lastVal := s.dataStore[s.cursor-1]
	s.dataStore = s.dataStore[:s.cursor-1]
	s.cursor--
	return lastVal, nil

}

// Push is a method implemented to append an element to the stack
func (s *Stack) Push(val interface{}) error {
	if s.IsFull() {
		return errors.New("Stack is full, there's no extra available spaces")
	}
	s.dataStore = append(s.dataStore, val)
	s.cursor++
	return nil
}

// IsFull is a method implemented to make sure if the stack is full
func (s *Stack) IsFull() bool {
	return s.cursor >= s.capSize
}

//IsEmpty is a method implemented to make sure if there's no any element in stack
func (s *Stack) IsEmpty() bool {
	return s.cursor == 0
}

// String is a method implemented to show stack as type string
func (s *Stack) String() string {
	return fmt.Sprint(s.dataStore)
}
