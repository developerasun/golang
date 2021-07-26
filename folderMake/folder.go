package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// Create a map that shows how many dates each month has.
// The map is not used in any other places, since it is only purpose of reminder.
// Do not inlcude this map in your program when you use this module.
var monDate = map[string]int{
	"1월":  31,
	"2월":  29,
	"3월":  31,
	"4월":  30,
	"5월":  31,
	"6월":  30,
	"7월":  31,
	"8월":  31,
	"9월":  30,
	"10월": 31,
	"11월": 30,
	"12월": 31,
}

func main() {

	// User enters a month to create folder.
	var userMonth int
	fmt.Println("폴더를 생성하고자 하는 월을 입력하세요:  ")
	fmt.Scanln(&userMonth)

	// monthDate function returns result variable and error.
	result, _ := monthDate(userMonth)

	// if result == 0, print program finish message.
	if result == 0 {
		fmt.Println("프로그램을 종료합니다.")

		// if result != 0, create folder.
	} else {
		for i := 1; i <= result; i++ {
			basepath := "./"
			baseName := "2021년" + strconv.Itoa(userMonth) + "월" + strconv.Itoa(i) + "일"
			folderPath := filepath.Join(basepath, baseName)

			_, err := os.Stat("test")
			if os.IsNotExist(err) {
				err := os.Mkdir(folderPath, 0755)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		// Print folder generation complete message when done.
		fmt.Printf("%d월 폴더 생성이 완료되었습니다.", userMonth)
	}
}

// Enter a month to decide how many dates you create in a folder
func monthDate(userMonth int) (result int, err error) {
	switch userMonth {
	case 1:
		return 31, nil
	case 2:
		return 28, nil // Ignore a leap month
	case 3:
		return 31, nil
	case 4:
		return 30, nil
	case 5:
		return 31, nil
	case 6:
		return 30, nil
	case 7:
		return 31, nil
	case 8:
		return 31, nil
	case 9:
		return 30, nil
	case 10:
		return 31, nil
	case 11:
		return 30, nil
	case 12:
		return 31, nil
	default:
		// Error handling when user enters invalid input data.
		fmt.Println("올바르지 않은 입력입니다.")
		return 0, err
	}
}
