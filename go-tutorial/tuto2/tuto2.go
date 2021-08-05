package main

import (
	"errors"
	"fmt"
)

var ErrInvalidNode = errors.New("Node is not vaild")

type Node interface {
	SetValue(v int) error
	GetValue() int
}

func createSLLNode(v int) Node {
	sn := NewSLLNode()
	sn.SetValue(v)
	return sn
}

func createPowerNode(v int) Node {
	sn := NewPowerNode()
	sn.SetValue(v)
	return sn
}

func NewSLLNode() *SLLNode {
	return &SLLNode{SNodeMessage: "This is a message from the normal Node"}
}

func NewPowerNode() *PowerNode {
	return &PowerNode{PNodeMessage: "This is a message from the power node"}
}

type SLLNode struct {
	next         *SLLNode
	value        int
	SNodeMessage string
}

// Receive represent a type implementing a method. In this case,
// *SLLNode is a data type that implements method SetValue.
// From Object-Oriented Programming perspective, receiver in golang is the same
// with class type(object type)
func (sNode *SLLNode) SetValue(v int) error {
	if sNode == nil {
		return ErrInvalidNode
	}
	sNode.value = v
	return nil
}

func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

type PowerNode struct {
	next         *PowerNode
	value        int
	PNodeMessage string
}

// The receiver of a method is allowed to be nil in golang
func (sNode *PowerNode) SetValue(v int) error {

	// Allowing nil type in receiver is a powerful feature in golang since it
	// enables developers to simply add validation checking in method like below
	// without having to specifying if it is nil whenever calling method in main function.
	if sNode == nil {
		return ErrInvalidNode
	}

	sNode.value = v + 10
	return nil
}

func (sNode *PowerNode) GetValue() int {
	return sNode.value
}

func main() {
	fmt.Println("============= Pointer PowerNode: Nil type receiver =============")
	// Poiter *PowerNode is not initialized. When not initiazlied, its zero value is nil
	var powernode *PowerNode
	fmt.Println(powernode.SetValue(5))

	fmt.Println()
	fmt.Println("============= Type switch in golang : interface  =============")
	mySLL := createSLLNode(5)
	myPower := createPowerNode(5)

	var mySlice []interface{}
	mySlice = append(mySlice, myPower, mySLL)

	// Type switch in golang: Since concreteN is interface type, it can store *SLLNode or *PowerNode.
	// Depending on your code, below switch ~ case statement will choose one of them differently.
	switch concreteN := mySlice[0].(type) {
	case *SLLNode:
		fmt.Println("Type is SLLNode, message: ", concreteN.SNodeMessage)
	case *PowerNode:
		fmt.Println("Type is PowerNode, message: ", concreteN.PNodeMessage)
	}

	fmt.Println()

	fmt.Println("============= Interface value assignment in golang  =============")
	var myInter Node

	// Create a deferred function to recover a program from goroutine panic
	// Be aware that deferred function should be ahead of error invoking. Otherwise,
	// program will just be terminated and then the deferred function will not execute.
	defer func() error {
		r := recover()
		fmt.Println(r, "Panic recovered!")

		// Value assignment into interface.
		myInter = &PowerNode{}
		fmt.Println(myInter.SetValue(5))
		return myInter.SetValue(5)
	}()

	fmt.Println(myInter.SetValue(5)) // this will invoke error since interface needs a value to be assigned

}
