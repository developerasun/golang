package main

import (
	"fmt"
	"sync"
)

func main() {
	var a, b = 10, 5
	var result int
	var wait sync.WaitGroup

	wait.Add(1)
	go func() {
		defer wait.Done()
		result = a + b
	}()

	wait.Wait()
	fmt.Printf("두수의 합: %d", result)
}
