package main

import "fmt"

func main() {
	// Be aware of underlying slice's memory
	test1 := memoryLeak()

	// Notice that the last element, 6, is still there.
	fmt.Println("My test1 slice here: ", test1, "Slice capacity of the test1 slice: ", test1[:cap(test1)])

	test2 := memoryNotLeak()

	// Notice that the last element, 6, is not there anymore since garbage collector has unassigned this.
	fmt.Println("My test2 slice here: ", test2, "Slice capacity of the test2 slice: ", test2[:cap(test2)])

	test3 := displayUnpacking()
	fmt.Println("My test3 slice here: ", test3, "Slice capacity of the test3 slice: ", test3[:cap(test3)])
}

// Create a function that garbage collector does not unassign
func memoryLeak() []int {
	s := []int{1, 2, 3, 4, 5, 6}
	return s[2:5]
}

// Create a function that garbage collector unassigns
func memoryNotLeak() []int {
	s := []int{1, 2, 3, 4, 5, 6}
	slice := make([]int, 3)
	copy(slice, s[2:5])

	return slice
}

// Create a function that shows unpacking in golang
func displayUnpacking() []int {

	// Create a slice using make built-in function
	mySlice := make([]int, 6)
	fmt.Println("Slice value before appending: ", mySlice)

	addSixHere := []int{1, 2, 3, 4, 5}

	// (your_slice)... : this is the same concept unpacking in Python
	result := append(mySlice, append(addSixHere, 6)...)
	fmt.Println("Slice value after appending: ", result)

	return result
}

// Basic concept of copy built-in function
// s1 := []int{1, 2, 3, 4, 5, 6}
// 	fmt.Println("s1: ", s1)

// 	s2 := make([]int, 3)

// 	// The copy built-in function copies elements from a source slice
// 	// into a destination slice. It returns a number of copied elements
// 	n := copy(s2, s1[2:4])

// 	fmt.Println("Number of copied elements: ", n)

// 	// Check S2 elements
// 	fmt.Println(s2)

// 	// Update the first element of slice s2. It does not affect s1 element now
// 	// since it is copied
// 	s2[0] = 333

// 	// Slice S1 has not been changed since now S1 and S2 is not the same slice.
// 	fmt.Println("S1: ", s1)
// 	fmt.Println("S2: ", s2)
