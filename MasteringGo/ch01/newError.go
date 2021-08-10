package main

import (
	"errors"
	"fmt"
)

// error in go is a data type unlike other programming language and
// each error has a different significance in terms of program termination.
func returnError(a, b int) error {
	if a == b {

		// errors.New takes a string argument and return error.
		// errors.Error functions compares error variable with string one
		// errors.New <-> errors.Error
		err := errors.New("Error in returnError() function!")
		return err
	} else {
		return nil
	}
}

func main() {
	err := returnError(1, 2)

	// Most of time, checking an error manually is inevitable in Go.
	if err == nil {
		fmt.Println("returnError() ended normally!")
		// calling os.Exit(integer that is not zero) from a function other than
		// main() function is a bad practice.
	} else {
		fmt.Println(err)
	}

	err = returnError(10, 10)
	if err == nil {
		// error handling -> related to Go codes handling error conditions
		// printing error output -> writing to standard error file descriptor
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}

	// errors.Error functions compares error variable with string one
	if err.Error() == "Error in returnError() function!" {
		fmt.Println("!!")
	}
}
