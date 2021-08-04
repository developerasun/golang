package main

import "fmt"

func main() {
	myX := 5

	// Function ChangeX takes myX's memory address as argument
	ChangeX(&myX)
	fmt.Println("The value of X is: ", myX)
}

// Referencing variable X to see a value
func ChangeX(X *int) {
	// Dereferencing X to get the value and update it.
	*X = 10
}
