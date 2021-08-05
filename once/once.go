package main

import (
	"fmt"
	"sync"
)

func main() {

	// Once is an object that will perform exactly one action.
	// A Once must not be copied after first use.
	var once sync.Once
	onceBody := func() {
		fmt.Println("Run only once")
	}

	done := make(chan bool, 2)

	// <-done
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()

	}

	for i := 0; i < 10; i++ {
		<-done
	}

}
