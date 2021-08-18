package main

import (
	"container/list"
	"fmt"
)

func main() {

	// List represents a doubly linked list. The zero value for List is an empty list ready to use.
	var x list.List

	// PushBack inserts a new element e with value v at the back of list l and returns e.
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	// Front returns the first element of list l or nil if the list is empty.
	// Next returns the next list element or nil.
	for e := x.Front(); e != nil; e = e.Next() {

		// The value stored with this element - since e.value's data type is interface,
		// type assertion needs to be done to remove ambiguity
		fmt.Println(e.Value.(int)) // print doubly list element in order

	}

}
