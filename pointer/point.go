package main

import "fmt"

func main() {
	myX := 5

	// Function ChangeX takes myX's memory address as argument
	ChangeX(&myX)
	fmt.Println("The value of X is: ", myX)

	// Unlike pointer in C, pointer in Go does have some more constraints

	a := int(5) // type assertion
	p := &a
	// p++ : No mathematic operations allowed on pointer.
	fmt.Print(p)

	floatA := 5.0
	floatP := &floatA
	// floatP = &a : Cannot convert between different types of pointers
	fmt.Print(floatP)

	// fmt.Print(p == floatP) : Cannot compare between different types of pointers

	// In unsafe package there is a pointer defined, which can bypass type system check.
	// Most of time, use of unsafe pointer is not recommended since as its name suggests, it is unsafe.

	// Two important features of unsafe package : 1) any pointer <-> unsafe.Pointer 2) unsafe.Pointer <-> uintptr
	// On pointer level, no mathematical operations supported(in golang)
	// Use uintptr if pointer-wise mathematical operations needed.
}

// Referencing variable X to see a value
func ChangeX(X *int) {
	// Dereferencing X to get the value and update it.
	*X = 10
}
