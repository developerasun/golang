package main

import "fmt"

// Treat interface a container. It can contain any data type and also be used
// as argument in fucntion.
func printVal(i interface{}) {
	fmt.Println(i)
}

func main() {
	// Interface is a dynamic type.
	var x interface{}
	var y interface{}
	// Interface type as an array.
	x = [3][2][2]int{}
	printVal(x)

	// Interface type as a string
	x = "test"
	printVal(x)

	// Interface type as a int
	x = 55
	y = 22
	y = y.(int)
	fmt.Printf("x의 타입: %T x값: %d, y의 타입: %T y값: %d", x, x, y, y)

}
