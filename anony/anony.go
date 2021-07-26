package main

import (
	"fmt"
)

func main() {

	// Create an anonymous function inside of main function.
	func() {
		fmt.Println("hello")
	}()

	// You can call the function right next to curly braces.
	func(a, b int) {
		result := a + b
		fmt.Println(result)
	}(2, 8)

	// You store a return value of anonymous fuction in a declared variable.
	result := func(a, b string) string {
		return a + b
	}("hello ", "world")
	fmt.Println(result)

	// Anonymous function is used to achieve memory efficiency.
	i, j := 20.4, 10.2
	divide := func(a, b float64) float64 {
		return a / b
	}(i, j)
	fmt.Println(divide)

}
