package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func hello(n int, w *sync.WaitGroup) {
	defer w.Done()

	r := rand.Intn(3)

	time.Sleep(time.Duration(r) * time.Second)
	fmt.Println(n)
}

func main() {
	wait := new(sync.WaitGroup)
	// Add 100 goroutines.
	wait.Add(100)

	// Create 100 goroutines.
	for i := 0; i < 100; i++ {
		go hello(i, wait)
	}

	// Wait till gorouine ends.
	wait.Wait()
}
