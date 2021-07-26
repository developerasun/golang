package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// runtime.NumCPU() returns a number of CPU of your computer.
	// runtime.GOMAXPROCS set how many CPUs you are going to use.
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println(runtime.GOMAXPROCS(0))

	var wait sync.WaitGroup
	wait.Add(100)

	for i := 0; i < 100; i++ {
		go func(n int) {
			defer wait.Done()
			fmt.Println(n)
		}(i)
	}

	wait.Wait()
}
