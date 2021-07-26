package main

import "fmt"

const (
	// iota starts with value 0 and increase ++1 each time. It is similar to enum function in other languages.
	// Use iota to increase code readability by assigning predefined numbers to const.
	Sunday    = 1 << iota
	Monday    // 1 << 1
	Tuesday   // 1 << 2
	Wednesday // 1 << 3
	Thursday  // 1 << 4
	Friday    // 1 << 5
	Saturday  // 1 << 6
)

func main() {
	stat := Tuesday
	fmt.Println(stat)
}
