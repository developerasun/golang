package main

import "fmt"

// Module 1 - Consumer info - adult, pregnant woman, youth
// Types of consumer info - 1) age 2) sex 3) pregnancy
// 1) age : 1 ~ 19 - teenager // 20 ~ 60 adult // 61 ~ 90 senior
// 2) sex : man // woman
// 3) pregnancy : T/F
// Function: 1) InfoEnter 2) Recommended amount 3) Actual amount 4) Compare amounts

// Create a map to index coffee drinks and corresponding caffeine amount.
var Drinks = map[string]float64{
	"LetsBe":  20,
	"Georgia": 10,
	"Maxim":   30,
}

// Create an abstract class : CaffeineCheck
type CaffeineCheck interface {
	Calculate(float64) float64
	ChooseCoffee(string)
}

// Create a embodying class : Coffee
type Coffee struct {
	product    string
	caffAmount float64
}

// Calculate total caffeine amount
func (c *Coffee) Calculate(number float64) float64 {
	total := c.caffAmount * number
	return total
}

// Choose which coffee drink to have
func (c *Coffee) ChooseCoffee(name string) {

	c.product = name

	if c.product == "LetsBe" {
		c.caffAmount += Drinks["LetsBe"]

	} else if c.product == "Georgia" {
		c.caffAmount += Drinks["Georgia"]

	} else {
		c.caffAmount += Drinks["Maxim"]
	}

}

func main() {
	// Declare a Coffee strcuture variable
	var myCaffeine Coffee

	// Choose drinks from the map
	myCaffeine.ChooseCoffee("Georgia")

	// Calculate total caffeine amount
	result := myCaffeine.Calculate(5)

	fmt.Println(result)

}

// Interface:
// figure A - var myInterface interface, myInterface = your_data_type
// figure B - func myInterface (i interface){i.your_interface_method} ===> Use this function in main function.
// figrue C - func (y your_struct) myInterface(i interface){i.your_interface_method} ===> declare a new struct variable
// and use interface functionaliy from the struct.

// 1. Create a functionA connected to the interface. This fucntion takes an embodying class(struct) as argument.
// 2. Depending on the struct, it categorizes which one is which.
// 3. Turn the functionA into a method connected to a new structA(embodying class).
// 4. Now you can declare a new variable for the new structA and use all the interface functionality.

// type User struct {
// 	name   string
// 	age    int
// 	sex    string
// 	weight float64
// 	isPreg bool
// 	info []string
// }

// func (u *User) EnterInfo(name, sex string, age int, weight float64, isPreg bool) {
// 	u.name = name
// 	u.sex = sex
// 	u.age = age
// 	u.isPreg = isPreg
// 	u.weight = weight
// }

// func (u *User) DecideCaffeineCheck() (amount float64) {
// 	if u.age <= 19 {
// 		amount = u.weight * 2.5
// 		return amount

// 	} else if u.age <= 60 && u.isPreg { // default bool value: false
// 		amount = 400
// 		return amount

// 	} else {
// 		amount = 300
// 		return amount
// 	}

// }
