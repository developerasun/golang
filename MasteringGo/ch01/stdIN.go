package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// the use of var keyword: 1) declaring a global variable(most of time)
	// 2) setting no initial value
	var f *os.File

	// three ways of getting user input: 1) cli args 2) asking user for input 3) read external file
	// os package's most common functionality is to use cli args.
	// os package is platform-independent(works well both in UNIX and MS Windows)
	f = os.Stdin
	defer f.Close()

	// Declare a new scanner with bufio package
	scanner := bufio.NewScanner(f)

	// Scan() function reads line by line from os.Stdin
	// stdIN, stdOUT are useful in UNIX pipes.
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())

		// Add loop escape condition
		if scanner.Text() == "stop" {
			break
		}

	}
}
