package main

import "fmt"

// Create an interface - abstract class
type Animal interface {
	Sound()
}

// Create a dog structure - embodying class
type Dog struct{}

// Create a dog's Sound method - embodying method
func (d *Dog) Sound() {
	fmt.Println("bark bark")
}

// Create a Cat structure - embodying class
type Cat struct{}

// Create a Cat's Sound method - embodying method
func (c *Cat) Sound() {
	fmt.Println("meow meow")
}

// Create a Cow structure.
type Cow struct{}

// Since the Cow structure does not have Sound() method, it is not
// an embodying class of the abstract class, Animal interface.
func (c *Cow) CowSound() {
	fmt.Println("ummm ummm")
}

func main() {

	// Declare a slice that has Animal as a data type.
	myArray := []Animal{}

	// Add Dog, Cat structures into the slice.
	myArray = append(myArray, &Dog{})
	myArray = append(myArray, &Cat{})

	// Below, you can't add structure Cow since it is missing method Sound()
	// so it can't be used a value of interface Animal.
	// myArray = append(myArray, &Cow{})

	// Print out index and value of the array.
	for index, val := range myArray {
		fmt.Println(index)
		fmt.Println(val)
	}

}
