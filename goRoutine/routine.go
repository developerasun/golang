package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(n int) {
	r := rand.Intn(3)
	// Convert r into time.Duration data type to caculate.
	time.Sleep(time.Duration(r) * time.Second)
	fmt.Println(n)
}

func main() {
	for i := 0; i < 100; i++ {
		go hello(i)
	}

	fmt.Scanln()
}
