package main

import "fmt"

func eat() {
	var sushi = "sushi"
	fmt.Println(sushi)
}

func main() {
	eat()
	// sushi = "not sushi" : undefined variable error (same with Javascript, local scope)

	for i := 0; i < 3; i++ {
		var bread = "bread"
		fmt.Println(bread)
	}
	// bread = "not bread" : undefined variable error (different from Javascript, Javascript
	// can access to the variable bread since it is outside 'function')

}