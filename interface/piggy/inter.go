package main

import "fmt"

// Declare an interface: a collection of method

type InsertMoney interface {
	Deposit() string
}

type Piggy interface {
	CountMoney() InsertMoney
}

type PiggyBank struct {
	money string
}

// Create three structures that belong to Piggy interface
type Won struct{}
type Dollar struct{}
type Euro struct{}

// Create three structures that belong to InsertMoney interface
type Counted_Won struct{}
type Counted_Dollar struct{}
type Counted_Euro struct{}

// Declare each method that is connected to Piggy structure.
func (p *PiggyBank) CountPiggy(piggy Piggy) {
	countedMoney := piggy.CountMoney()
	p.money += countedMoney.Deposit()
}

// Declare an executing function.
func (p *PiggyBank) Deposit() string {
	return "Open " + p.money + "Close"
}

// Declare corresponding methods: Won, Dollar, Euro
// AND Counted_Won, Counted_Dollar, Counted_Euro
func (w *Won) CountMoney() InsertMoney {
	return &Counted_Won{}
}

func (c *Counted_Won) Deposit() string {
	return " + Won"
}

func (w *Dollar) CountMoney() InsertMoney {
	return &Counted_Dollar{}
}

func (c *Counted_Dollar) Deposit() string {
	return " + Dollar"
}

func (w *Euro) CountMoney() InsertMoney {
	return &Counted_Euro{}
}

func (c *Counted_Euro) Deposit() string {
	return " + Euro"
}

func main() {
	myPiggy := &PiggyBank{}
	won := &Won{}
	dollar := &Dollar{}
	euro := &Euro{}

	fmt.Println("My Piggy: ")
	myPiggy.CountPiggy(won)
	myPiggy.CountPiggy(dollar)
	myPiggy.CountPiggy(euro)
	fmt.Println(myPiggy)
}
