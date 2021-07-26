package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	wait.Add(102)

	str := "goorm"

	go func() {
		defer wait.Done()
		fmt.Println("Hello")
	}()

	go func() {
		defer wait.Done()
		fmt.Println(str)
	}()

	for i := 0; i < 100; i++ {
		go func(n int) {
			defer wait.Done()
			fmt.Println(n)
		}(i)
	}
	wait.Wait()
}

// WaitGroup order
// 1. wait.Add(): Set the number of goroutines
// 2. wait.Done(): Notice goroutine function call finishes
// 3. wait.Wait(): Wait all the goroutine functions finish.
