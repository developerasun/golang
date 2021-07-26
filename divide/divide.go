package main

import (
	"errors"
	"fmt"
)

// Declare user-defined error
var zeroErr error = errors.New("0으로 나눌 수 없습니다.")

func divide(a, b float32) (result float32, err error) {
	if b == 0 {
		return 0, zeroErr
	}
	result = a / b
	return
}

func main() {
	var num1, num2 float32
	fmt.Scanln(&num1, &num2)

	result, err := divide(num1, num2)

	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
