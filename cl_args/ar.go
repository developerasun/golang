package main

import (
	"fmt"
	"os"
)

// Create a function that takes a string array as argument.
func cmdArgs(cmd ...string) []string {
	return cmd
}

func main() {

	// Command-line arguments are a way to provide the parameters or arguments to the main function of a program.
	// Similarly, In Go, we use this technique to pass the arguments at the run time of a program.
	// arguments := os.Args

	// The first argument of os.Args is always program name : program name with full path(Os Filepath Output)
	// os.Args is a string array.
	myProgramName := os.Args[0]
	fmt.Println(myProgramName)

	myArguments := cmdArgs(os.Args...)
	// Print all except Os.Filepath output
	fmt.Println(myArguments[1:])

}
