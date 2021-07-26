package main

import (
	"fmt"
	"math"
)

// Struct data type is a collection of user-custom data.
// Interface is a collection of user-custom method.

type geometry interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r1 := Rect{10, 20}
	c1 := Circle{10}
	r2 := Rect{12, 14}
	c2 := Circle{5}

	// func printMeasure takes structs as arguments, because
	// method area and perimeter is a method of struct Rect and Circle.
	printMeasure(r1, c1, r2, c2)
}

func printMeasure(m ...geometry) {
	for _, val := range m {
		fmt.Println(val)
		fmt.Println(val.area())
		fmt.Println(val.perimeter())
	}
	fmt.Println(len(m))
}
