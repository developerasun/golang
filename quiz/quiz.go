package main

import "fmt"

func main() {

	var a, b int = 26, 81

	if check := rain80(a, b); check {
		fmt.Println("덥고 비가 옵니다.")
	} else if check2 := rain20(a, b); check2 {
		fmt.Println("덥고 습합니다.")
	} else if check3 := lessRain20(a, b); check3 {
		fmt.Println("야외 활동하기 좋습니다.")
	} else {
		if check4 := notTemp25(a, b); check4 {
			fmt.Println("야외 활동하기 좋지 않습니다.")
		} else {
			fmt.Println("좋은 날씨입니다.")
		}
	}
}

func rain80(a, b int) bool {
	check1 := a >= 25 && b >= 80
	return check1
}

func rain20(a, b int) bool {
	check2 := a >= 25 && b >= 20
	return check2
}

func lessRain20(a, b int) bool {
	check3 := a >= 25 && b < 20
	return check3
}

func notTemp25(a, b int) bool {
	check4 := a < 25 && (a < 10 || b >= 80)
	return check4
}
