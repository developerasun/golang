package main

import (
	"github.com/designerasun/golang/defer/timetravel"
)

func main() {

	var MyController timetravel.RemoteControl
	var Jonghyun timetravel.Person
	MyController = Jonghyun

	MyController.Press().CheckState()
	MyController.Press().Break()
	MyController.Press().InitTimeMachine()

}
