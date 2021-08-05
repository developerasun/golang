package main

import "fmt"

func main() {
	printType("Good bye World")
	printType(55678)
	printType(99.99)
}

// function printType can take any data type. Since it takes an empty interface.
// This style of coding makes your code very flexible.
func printType(i interface{}) {
	switch i := i.(type) {
	case string:
		fmt.Println("This is a string type", i)
	case int:
		fmt.Println("This is a int type", i)
	case float64:
		fmt.Println("This is a float64 type", i)
	}
}
