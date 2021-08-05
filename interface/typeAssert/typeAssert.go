package main

import (
	"fmt"
	"log"
)

func main() {

	defer func() {
		r := recover()
		log.Println("Error occured and recovered: ", r)
	}()

	// Type Assertion in golang: used to remove ambiguity from interface variable
	// Extracting actual data type value help by interface variable.
	// Type assertion syntax : t := value.(typeName)

	// Declare an interface that has string value
	var value interface{} = "IamInterface"

	// value.(string) -> concrete type we want to check
	var value1 string = value.(string)
	fmt.Println("Printing the concrete value1: ", value1)

	// Type assertion can return two values: 1) concrete value 2) boolean(assertion check)
	var floatValue interface{} = 99.99
	myValue, found := floatValue.(float64)
	if found {
		fmt.Println("float value found!: ", myValue)

	} else {
		fmt.Println("float value not found!")
	}

	// interface conversion error : interface {} is string, not int
	var value2 int = value.(int)
	fmt.Println("Printing the concrete value: ", value2)

}
