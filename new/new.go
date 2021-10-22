package main

import "fmt"

// new(Type) *Type
// The new built-in function allocates memory. The first argument is a type,
// not a value, and the value returned is a pointer to a newly
// allocated zero value of that type.
// Unlike new built-in function, make built-in function can take only
// slice, map, and channel as argument.
// Both new and make allocates and initialize memory

func main() {
	result := make(map[int]string)
	fmt.Println(result)
}
