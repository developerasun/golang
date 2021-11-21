package main

import (
	"fmt"
	"unsafe"
)

// In unsafe package there is a pointer defined, which can bypass type system check.
// Most of time, use of unsafe pointer is not recommended since as its name suggests, it is unsafe.

// Two important features of unsafe package : 1) any pointer <-> unsafe.Pointer 2) unsafe.Pointer <-> uintptr
// On pointer level, no mathematical operations supported(in golang)
// Use uintptr if pointer-wise mathematical operations needed.
func main() {
	// Declare struct variable
	var p Person

	// Initialize object
	p = Person{name: "Jonghyun", age: 27}

	//
	name := (*string)(unsafe.Pointer(&p))
	*name = "Not Jonghyun"

	// With Offsetof(), the position of each member in a struct can be found out
	// and their memory can be accessed and updated accordingly.
	age := (*int)(unsafe.Pointer((uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.age))))
	*age = 277

	// Check the changed value
	fmt.Println("age: ", p.age, "name: ", p.name)
}

type Person struct {
	name string
	age  int
}
