package main

import (
	"fmt"
)

var (
	balance int
)

func init() {
	balance = 1000
}

func main() {

	go func(value int) {

		for {
			balance += value
			fmt.Printf("Current balance1: %d\n", balance)
		}

	}(10)

	go func(value int) {

		for {
			balance -= value
			fmt.Printf("Current balance2: %d\n", balance)
		}

	}(10)

	fmt.Scanln()
	fmt.Printf("Latest balance: %d", balance)

}
