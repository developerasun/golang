package main

import "fmt"

func main() {
	var str = "Hello there!"
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(str, i)
		}
		done <- true
	}()
	<-done
}
