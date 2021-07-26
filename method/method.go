package main

import "fmt"

// Declare user-customed data type struct.
type account struct {
	balance   int
	firstName string
	lastName  string
}

// Create a pointer-based method.
func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

// Create a value-based method.
func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

// Create a valued-based method with return value.
func (a3 account) withdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}

func main() {
	var mainA *account = &account{100, "Bruce", "Lee"}

	// Use a pointer-based method.
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance) // result: balance = 70

	// Use a value-based method.
	mainA.withdrawValue(20)
	fmt.Println(mainA.balance) // result: balance = 70

	// Use a value-based method with return value
	var mainB account = mainA.withdrawReturnValue(20)
	fmt.Println(mainB.balance, mainA.balance) // result: balance = 50

	// Use a pointer-based method.
	mainB.withdrawPointer(30)
	fmt.Println(mainB.balance) // result: balance = 20

	

}
