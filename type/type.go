package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	var myVar1 = "hello"
	var myVar2 = time.Second
	var myVar3 = 34.3
	myType(myVar1, myVar2, myVar3)
}

func myType(data1 string, data2 time.Duration, data3 float64) {

	// reflect.TypeOf function returns type of data(parameter)
	fmt.Println(reflect.TypeOf(data1))
	fmt.Println(reflect.TypeOf(data2))
	fmt.Println(reflect.TypeOf(data3))
}
