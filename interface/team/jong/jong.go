package main

// Code written by Jonghyun Sung
import "fmt"

type Restaurant struct {
	DongDaeguChinese
	DongDaeguJapan
	DongDaeguKorean
}

type house struct {
}

// Create a restuarant 1
type DongDaeguKorean struct {
	kimchi  string
	bulgogi string
}

// Createa a restaurant 2
type DongDaeguJapan struct {
	sushi string
	ramen string
}

// Create a restaurant 3
type DongDaeguChinese struct {
	dimsum string
	tofu   string
}

// Declare an interface
type Recipe interface {
	waitingTime() string
	deliver() string
}

// Create a method for DongDaeguKorean
func (dk *DongDaeguKorean) waitingTime(time int) (serving string) {
	if time < 5 {
		dk.kimchi = "잘 익은 김치"
		return dk.kimchi
	} else {
		dk.bulgogi = "달달한 불고기"
		return dk.bulgogi
	}
}

// Create a method for DongDaeguJapan
func (dj *DongDaeguJapan) waitingTime(time int) (serving string) {
	if time < 5 {
		dj.sushi = "연어초밥"
		return dj.sushi
	} else {
		dj.ramen = "달달한 라면"
		return dj.ramen
	}
}

// Create a method for DongDaeguChinese
func (dc *DongDaeguChinese) waitingTime(time int) (serving string) {
	if time < 5 {
		dc.dimsum = "쫄깃한 딤섬"
		return dc.dimsum
	} else {
		dc.tofu = "부드러운 두부"
		return dc.tofu
	}
}

func (h *house) deliver(addr string) {
	fmt.Printf("%s로 배달이 완료되었습니다.", addr)
}

func main() {

	var wait int
	var addr string

	whereEat := Restaurant{}
	Gyeongsan := house{}

	fmt.Println("고블록 식당에 오신걸 환영합니다. 메뉴는 대기 시간에 따라 달라집니다. 예상 대기 시간을 정하십시오: ")
	fmt.Scanln(&wait)
	fmt.Println("집 주소를 입력하십시오: ")
	fmt.Scanln(&addr)

	myDish := whereEat.DongDaeguChinese.waitingTime(wait)
	fmt.Println(myDish)
	Gyeongsan.deliver(addr)

}
