package main

import "fmt"

func panicTest() {
	defer func() {
		r := recover()
		fmt.Println(r)
	}()

	var a = [4]int{1, 2, 3, 4}

	for i := 0; i < 10; i++ {
		fmt.Println(a[i])
	}
}

func main() {
	panicTest()
	fmt.Println("Hello")
}
