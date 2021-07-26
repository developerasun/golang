package main

import "fmt"

func main() {
	num := []int{1, 2, 3}
	student := []string{"jake", "nana", "chris"}
	bools := []bool{true, false, false}

	fmt.Println(num, student, bools)

	// Use append function to add new value to your array.
	num = append(num, 4)
	fmt.Println(num)

	// You can add several values at once.
	num = append(num, 5, 6, 7)
	fmt.Println(num, num[6])

	// Slice example: newNum[2:5] -> select 2nd, 3rd, 4th element
	var newNum []int = num[2:5]
	fmt.Println(newNum)

	// You can slice literal.
	var s string = "Hello world"
	fmt.Println(s[5:8])

	// Declare array in go: 1) shorthanded command 2) var keyword
	myArray1 := []int{1, 2, 3}
	var myArray2 [2]string
	myArray2[0] = "hello"
	myArray2[1] = "world"

	fmt.Println(myArray1, myArray2)

	// Create a map that has: string as key, int as value
	myMap := make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3
	fmt.Println(myMap)

	myMap2 := map[string]int{
		"index1": 1,
		"index2": 2,
		// Don't forget to end your map with , if you declare map in this way.
		"index3": 3,
	}
	fmt.Println(myMap2)

	// Declare pointer
	var a int
	var b *int

	a = 10
	b = &a

	fmt.Println(a, b)

	*b = 20
	fmt.Println(a)

	// Pointer cannot store value itself. Look at d = 10.
	// var c int
	// var d *int
	// a = 10
	// d = 10

	var myNum int = 1
	var numPtr *int = &myNum
	fmt.Println(&myNum, numPtr)

	// Initialize your pointer with new() function. new() allocates memory
	// and return pointer with nil value.
	var numPtr2 *string = new(string)
	fmt.Println(numPtr2)

}
