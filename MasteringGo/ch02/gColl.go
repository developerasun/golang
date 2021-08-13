package main

import (
	"fmt"
	"runtime"
	"time"
)

// Create a function to avoid repeated codes
func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}

// Garbagae collection(GC) is the process of freeing up memory that is not
// being used and it happens concurrently, meaning not before or after the
// executetion of the program.
func main() {
	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {

		// var s1 []int : var keyword delcares slice with nil value
		// make(your_slice, length, capacity)
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}
	printStats(mem)

	// The operation of the Go garbage collector is based on tricolor algorithm(tricolor mark-and-sweep algorithm)
	// The algorithm has an invariant: No objects in black set can point to those of white set.
	// Gray set acts like a buffer zone between white and black set
	// White set is the place where garbages are collected
	// New objects must go to gray set & when a pointer of the program is moved the object(the pointer points to) is colored gray.
	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
		time.Sleep(5 * time.Second)
	}
	printStats(mem)
}
