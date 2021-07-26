package main

import "fmt"

// Declare a Dog structure - child class
type Dog struct{}

// Declare the Dog's method1 - sound()
func (d *Dog) sound() {
	fmt.Println("bark bark")
}

// Declare the Dog's method2 - sleep()
func (d *Dog) sleep() {
	fmt.Println("Dog - sleep sleep")
}

// Declare the Dog's method3 - eat().
// This method is not executed since method eat() is not defined in interface.
func (d *Dog) eat() {
	fmt.Println("yummy yummy")
}

// Declare a Cat structure - child class
type Cat struct{}

// Declare the Cat's method1 - sound()
func (c *Cat) sound() {
	fmt.Println("meow meow")
}

// Declare the Cat's method2 - sleep()
func (c *Cat) sleep() {
	fmt.Println("Cat - sleep sleep")
}

// Declare an interface - abstract class, parent class
type AnimalInspector interface {
	// Abstract method
	sound()
	sleep()
}

// Create a function that is connected to the inferface.
// This function take an interface as argument outside. But actually, it takes
// child class and execute child class's methods.
func myAnimal(a AnimalInspector) {
	a.sound()
	a.sleep()
}

func main() {

	var myDog Dog
	var myCat Cat

	myAnimal(&myDog)
	myAnimal(&myCat)

}
