package main

import "fmt"

func hello() {
	fmt.Println("hello")
}

func world() {
	fmt.Println("world")
}

func main() {
	// var a, b int = 10, 0
	// defer fmt.Println("done")

	// result := a / b
	// fmt.Println(result)

	defer world()
	hello()

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}
