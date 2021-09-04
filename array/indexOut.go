package main

import "fmt"

func main() {
	myArr := []string{}
	myArr = append(myArr, "a", "b", "c")

	for i := 0; i < 4; i++ {
		fmt.Println(myArr[i])
	}
}
