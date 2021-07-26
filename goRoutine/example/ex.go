package main

import "fmt"

func main() {
	var num1, num2 int
	var c chan int

	fmt.Scanln(&num1, &num2)
	result := add(num1, num2, c)
	fmt.Println(result)
}

func add(num1, num2 int, c chan int) <-chan int {
	c <- num1 + num2
	return c
}
