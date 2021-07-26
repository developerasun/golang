package main

import "fmt"

func main() {
	var nums = []int{10, 12, 13, 14, 16}

	addAnonymous := func(nums ...int) (result int) {
		for i := 0; i < len(nums); i++ {
			result += nums[i]
		}
		return
	}
	fmt.Println()
	fmt.Println(addAnonymous(nums...))
	fmt.Println(addDeclared(nums...))

	add := func() {
		fmt.Println("익명 함수를 호출했습니다.")
	}

	add()

}

func addDeclared(nums ...int) (result int) {
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return result
}

func add() {
	fmt.Println("선언 함수를 호출했습니다")
}
