package main

import (
	"fmt"
	"math"
)

func main() {
	x := 0.0
	y := 1.0 / x
	z := x * y
	fmt.Println(z)

	tt := math.NaN()
	fmt.Println(tt)
}
