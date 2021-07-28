package main

import (
	"fmt"
	"sync"
)

var (
	// mutex   sync.Mutex
	balance int
)

func init() {
	balance = 1000
}

func main() {

	var wait sync.WaitGroup

	for counter := 0; counter < 500000; counter++ {

		wait.Add(2)

		go func(value int) {
			// mutex.Lock()
			balance += value
			// mutex.Unlock()
			fmt.Printf("=========Number of try: %d=========\n", counter)
			fmt.Printf("Current balance1: %d\n", balance)
			wait.Done()
		}(700)

		go func(value int) {
			// mutex.Lock()
			balance -= value
			// mutex.Unlock()
			fmt.Printf("=========Number of try: %d=========\n", counter)
			fmt.Printf("Current balance2: %d\n", balance)
			wait.Done()
		}(700)

		wait.Wait()
	}

	fmt.Printf("Latest balance: %d", balance)

}
