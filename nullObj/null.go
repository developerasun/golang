package main

import (
	"fmt"
	"log"
	"reflect"
)

type nullObj struct{
	message string
}

func (n nullObj) createNullObj() string {
	n.message = "I am null"
	return n.message
}

func main() {
	// var name interface{}
	var myError error
	var name = "jake"
	if (reflect.TypeOf("") != reflect.TypeOf(name)) {
		fmt.Println("error occured") // define what has to occur imperatively
		log.Fatal(myError)
	}

	fmt.Println(reflect.TypeOf(""))
	var check = reflect.TypeOf("jake")
	fmt.Println(reflect.Type(check))

	myNullObj := new(nullObj)
	
	var result = myNullObj.createNullObj()
	fmt.Println(result)

}