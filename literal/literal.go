package main

import "fmt"

func main() {
	// literals in go: 1) double quotation: "" 2) back quote: ``
	// back quote example ignores escape character
	rawLiteral := `new world\n
	new world\n
			new world`

	// double quotation example perceives escape character
	doubleLiteral := "\nnew world\n new world\n new world"

	// print out the two vars
	fmt.Println(rawLiteral, doubleLiteral)

}
