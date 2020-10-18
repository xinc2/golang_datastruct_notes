package arraylist

import (
	"errors"
	"fmt"
)

// List defines a general interface including some contents
type List interface {
	Size() int
	Get(index int) (interface{}, error)      // get the element by the index provided
	Set(index int, val interface{}) error    // modify element
	Insert(index int, val interface{}) error // insert an element
	Append(val interface{})                  // append an element at the end of list
	Clear()                                  // clear all elements in the list
	Delete(index int) error                  // delete an element
	String() string                          // display the list as string
	Iterator() Iterator                      // construct an interator
}

// ArrayList is data structure including int, string, etc
type ArrayList struct {
	dataStore []interface{}
	theSize   int
}

// NewArrayList is used to make an instance for ArrayList
func NewArrayList() *ArrayList {
	list := new(ArrayList)
	list.dataStore = make([]interface{}, 0, 10)
	list.theSize = 0
	return list
}

// checkListIsFull
func (al *ArrayList) checkListIsFull() {
	if al.theSize == cap(al.dataStore) {
		// double the memory assignment
		newDataStore := make([]interface{}, al.theSize, 2*al.theSize)
		// assign the original values to the new memory area
		for i := 0; i < al.theSize; i++ {
			newDataStore[i] = al.dataStore[i]
		}
		al.dataStore = newDataStore
	}
}

// Size is a method used to get the length of the list
func (al *ArrayList) Size() int {
	return al.theSize
}

// Get is a method used to obtain the expected element by the index
func (al *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index > al.theSize {
		return nil, errors.New("Index out of the bound")
	}
	return al.dataStore[index], nil
}

// Set is a method used to update an vlaue by the index provided
func (al *ArrayList) Set(index int, val interface{}) error {
	if index < 0 || index > al.theSize {
		return errors.New("Index out of the bound")
	}
	al.dataStore[index] = val
	return nil

}

// Insert a method used to insert an vlaue by the index provided
func (al *ArrayList) Insert(index int, val interface{}) error {
	if index < 0 || index > al.theSize {
		return errors.New("Index out of the bound")
	}
	// monitor memory, assign new memory if it's used up
	al.checkListIsFull()
	// assign new memory for the inserted val
	al.dataStore = al.dataStore[:al.theSize+1]
	// Starting from the index bit,
	// the memory location of all data moves one bit backward
	for i := al.theSize; i > index; i-- {
		al.dataStore[i] = al.dataStore[i-1]
	}
	al.dataStore[index] = val
	al.theSize++
	return nil
}

// Append is method used to add the val at the end of list
func (al *ArrayList) Append(val interface{}) {
	al.dataStore = append(al.dataStore, val)
	al.theSize++
}

// Clear is a method
func (al *ArrayList) Clear() {
	al.dataStore = make([]interface{}, 0, 10)
	al.theSize = 0
}

// Delete is a method used to delete an element by the index
func (al *ArrayList) Delete(index int) error {
	if index < 0 || index > al.theSize {
		return errors.New("Index out of the bound")
	}
	al.dataStore = append(al.dataStore[:index], al.dataStore[index+1:]...)
	al.theSize--
	return nil
}

//String is method used to show the list as the type of string
func (al *ArrayList) String() string {
	return fmt.Sprint(al.dataStore)
}
