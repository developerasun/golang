package main

import (
	"fmt"
	"math"
)

type geometry interface {
	surface_area() float64
	volume() float64
}

type cylinder struct {
	radius float64
	height float64
}

type rec_prism struct {
	side1, side2, side3 float64
}

// Declare a method for cylinder's surface_area
func (c cylinder) surface_area() float64 {
	return (math.Pi*c.radius*c.radius)*2 + (2*math.Pi*c.radius)*c.height
}

// Declare a method for cylinder's volume
func (c cylinder) volume() float64 {
	return math.Pi * c.radius * c.radius * c.height
}

// Declare a method for rectangular prism's surface area.
func (r rec_prism) surface_area() float64 {
	return 2*r.side1*r.side2 + 2*r.side2*r.side3 + 2*r.side3*r.side1
}

// Declare a method for rectangular prism's volume.
func (r rec_prism) volume() float64 {
	return r.side1 * r.side2 * r.side3
}

func main() {

	cy1 := cylinder{radius: 10, height: 10}
	cy2 := cylinder{radius: 4.2, height: 15.6}
	cu1 := rec_prism{10.5, 20.2, 20}
	cu2 := rec_prism{4, 10, 23}

	printMeasure(cy1, cy2, cu1, cu2)
}

func printMeasure(g ...geometry) {
	for _, val := range g {
		fmt.Printf("%.2f %.2f\n", val.surface_area(), val.volume())
	}
}
