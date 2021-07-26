package main

import "fmt"

// Declare a parrot interface that mimics Animal or Human.
// If an interface(A) has an interface(B) inside,
// the former one can call the latter's methods.
type Parrot interface {
	Animal
	Human
}

// Create a Animal abstract class - Animal
type Animal interface {
	Sound()
}

// Create a embodying class and embodying method - Dog
type Dog struct{}

func (d *Dog) Sound() {
	fmt.Println("bark bark")
}

// Create a embodying class and embodying method - Dog
type Cat struct{}

func (d *Cat) Sound() {
	fmt.Println("meow meow")
}

// Create a Animal abstract class - Human
type Human interface {
	Sound()
}

// Create a embodying class and embodying method - Man
type Man struct{}

func (m *Man) Sound() {
	fmt.Println("I'm a man")
}

// Create a embodying class and embodying method - Woman
type Woman struct{}

func (w *Woman) Sound() {
	fmt.Println("I'm a woman")
}

func main() {
	// Declare an interface variable. Interface is also one of the data types.
	var mimicker Parrot

	// If mimicker interface takes a Dog structure, it sounds like a dog.
	mimicker = &Dog{}
	mimicker.Sound()

	// If mimicker interface takes a Man structure, it sounds like a Man.
	mimicker = &Man{}
	mimicker.Sound()

}
