package main

import "fmt"

func main() {
	// Recover function restores panic error and let program re-run.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			// Use recursive function to call main function again.
			main()
		}
	}()

	var num1, num2 int
	fmt.Scanln(&num1, &num2)
	result := num1 / num2
	fmt.Println(result)

}
