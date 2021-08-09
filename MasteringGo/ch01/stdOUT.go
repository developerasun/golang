package main

import (
	"io"
	"os"
)

func main() {
	myString := ""

	// os package is used to read cli args.
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument!" // a hardcoded message
	} else {
		myString = arguments[1]
	}

	// stdOUT.go uses the io package instead of fmt package.
	// io.WriteString takes two args: 1) os.File(io.Writer) 2) string value
	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}
