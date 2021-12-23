package main

import "fmt"

const (
	// iota starts with value 0 and increase ++1 each time. It is similar to enum function in other languages.
	// Use iota to increase code readability by assigning predefined numbers to const.
	Sunday    = iota // 0
	Monday    // 1
	Tuesday   // 2
	Wednesday // 3
	Thursday  // 4
	Friday    // 5
	Saturday  // 6
)

func main() {
	stat := Saturday

	fmt.Println(stat)
}
