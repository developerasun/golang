package main

import "fmt"

// Golang does not provide class keyword and inheritance explicitly.
// Instead it replaces the terms as: class = struct + method // inheritance = embedding
// But here, we use the terms, class and inheritance, for our convenience.

// Create a Animal class - parent class.
type Animal struct {
	species string
	age     int
}

// Create a method PrintDog() connected to Animal - parent class' method.
func (a *Animal) PrintDog() {
	fmt.Println("Connected to Animal structure - Animal species: dog")
}

// Create a Dog1 class - child class.
type Dog1 struct {
	// Dog1 structure has the Animal structure as a field value in a composition manner.
	species Animal // (field name) (structure) - composition.
	fur     string
	bark    string
}

// Create a Dog2 class - child class.
type Dog2 struct {
	Animal
	fur  string
	bark string
}

// Create a Dog3 class - child class.
type Dog3 struct {
	// Dog3 structure has the Animal structure as a field value in a inheritance manner.
	Animal // (structure) - inheritance.
	fur    string
	bark   string
}

// Method Overriding - below PrintDog() method overrides parent class Animal's method PrintDog()
func (d *Dog3) PrintDog() {
	var x, y int
	x = 10
	y = 20
	fmt.Println(x + y)
	fmt.Println("========== Method overriding ==========")
	fmt.Println("Connected to Dog3 structure - Animal species: dog")
}

func main() {
	var myDog1 Dog1
	var myDog2 Dog2
	var myDog3 Dog3

	// myDog1 has a Animal class - composition
	myDog1.species.PrintDog()
	// myDog1.PrintDog() ===> Can't access it like this.

	// Below, it prints out Animal's method PrintDog()
	myDog2.PrintDog()

	// myDog3 is a Animal class - inheritance
	// Below, it prints out an overriden method PrintDog()
	myDog3.PrintDog()

}
