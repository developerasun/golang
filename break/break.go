package main

import "fmt"

func main() {

	for i := 1; i < 16; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}

}

// for & label & break practice
// i := 0
// TEST1:
// 	for {
// 		if i == 0 {
// 			break TEST1
// 		}
// 	}
// 	fmt.Println("End")

// for & break practice
// var sum = 0
// 	var i int
// 	i = 1

// 	for {
// 		sum += i
// 		if sum > 100 {
// 			break
// 		}
// 		i++
// 	}

// 	fmt.Println("1에서 ", i, " 까지 더하면 처음으로 100이 넘어요")
// 	fmt.Println("합계: ", sum)
