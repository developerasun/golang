package main

import (
	"fmt"
	"log"

	"github.com/designerasun/greetings"
)

func main() {
	// enter below command in terminal to import Hello function into your local path
	// go mod edit -replace "old path"="new path(your local path)"
	// command might be different by your terminal types: powershell, git bash ...

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)
	// non-nil value -> error value
	if err != nil {
		// if you get an error, print the error and stop the program.
		log.Fatal(err)
	}
	fmt.Println(messages)
	result := greetings.NewHi()
	fmt.Println(result)
}
