package main

import "fmt"

func main() {
	myFunc("Count up to 5", 1, 2, 3, 4, 5)
	area(8, 10)
	width, height := multiply(5, 6)
	fmt.Println(width, height)
}

func myFunc(s string, number ...int) {
	fmt.Println(s)
	for i := 0; i < len(number); i++ {
		fmt.Println(number[i])
	}

}

// Create a function to calcualte rectangle area
func area(w, h int) {
	for j := 1; j < h; j++ {
		result := w * j
		fmt.Printf("rectagle %d area: %d\n", j, result)
	}
}

func multiply(w, h int) (int, int) {
	return w * 2, h * 2
}
