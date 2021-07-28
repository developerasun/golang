package main

import (
	"github.com/golang/caffeine/drink"
	"github.com/golang/caffeine/journal"
)

// Daily caffeine intake checker: tracking how many drinks that contain
// caffeine one person takes and document it daily.

// Module 1 - Consumer info - adult, pregnant woman, youth
// Module 2 - Drink info - types of drink containing caffeine
// Module 3 - Caffeine journal showing statistics or graphs.

func main() {
	// consumer.EnterInfo()
	drink.DrinkType()
	journal.DrawGraph()

}
