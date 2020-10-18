package arraylist

import (
	"errors"
)

// Iterator is an interface iterator
type Iterator interface {
	HasNext() bool
	Next() (interface{}, error)
	Remove()
	GetIndex() int
}

// Iterable is an interface
type Iterable interface {
	Iterator() Iterator
}

// ArrayListIterator is a data structure
type ArrayListIterator struct {
	list   *ArrayList
	cursor int
}

// Iterator is a method used to construct an iterator
func (al *ArrayList) Iterator() Iterator {
	it := new(ArrayListIterator)
	it.cursor = 0
	it.list = al
	return it
}

// HasNext is a method used to make sure if there exist the next one
func (ali *ArrayListIterator) HasNext() bool {
	return ali.cursor < ali.list.theSize
}

// Next is a method used get the next one element
func (ali *ArrayListIterator) Next() (interface{}, error) {
	if !ali.HasNext() {
		return nil, errors.New("No next one element can be found")
	}
	val, err := ali.list.Get(ali.cursor)
	ali.cursor++
	return val, err
}

// Remove is a method
func (ali *ArrayListIterator) Remove() {
	ali.cursor--
	ali.list.Delete(ali.cursor)
}

// GetIndex is a method used to obtain the current index of the element
func (ali *ArrayListIterator) GetIndex() int {
	return ali.cursor
}
