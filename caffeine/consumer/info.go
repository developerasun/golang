package main

import "fmt"

// Module 1 - Consumer info - adult, pregnant woman, youth
// Types of consumer info - 1) age 2) sex 3) pregnancy
// 1) age : 1 ~ 19 - teenager // 20 ~ 60 adult // 61 ~ 90 senior
// 2) sex : man // woman
// 3) pregnancy : T/F

type User struct {
	name   string
	age    int
	sex    string
	weight float64
	isPreg bool
}

func (u *User) EnterInfo(name, sex string, age int, weight float64, isPreg bool) {
	u.name = name
	u.sex = sex
	u.age = age
	u.isPreg = isPreg
	u.weight = weight
}

func (u *User) DecideCaffeine() (amount float64) {
	if u.age <= 19 {
		amount = u.weight * 2.5
		return amount

	} else if u.age <= 60 && u.isPreg { // default bool value: false
		amount = 400
		return amount

	} else {
		amount = 300
		return amount
	}

}

func main() {

	fmt.Println("hello world")
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
