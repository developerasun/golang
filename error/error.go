package main

import "fmt"

func main() {
	var input string
	_, err := fmt.Scanln(&input)

	if err != nil {
		panic(err)
	}

	fmt.Println(input)
}
