package main

import "fmt"

func panicTest() {
	// Use recover function inside of deferred function
	// to avoid program termination. Remember that you should always
	// use defer & recover togehter to do error-handling.
	defer func() {
		r := recover()
		fmt.Println(r)
	}()

	var a = [4]int{1, 2, 3, 4}
	for i := 0; i < 10; i++ {
		fmt.Println(a[i])
	}
}

func main() {
	panicTest()
	fmt.Println("hello world")
}
