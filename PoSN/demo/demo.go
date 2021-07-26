package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

func main() {

	// Create a safe and safeNumber.
	mySafe := safe{}
	mySafeNumber := safePassword{}

	// Enter a user name, NFT token ID, and locktime to store.
	// Exmaple - owner: jonghyun, lock_time: 20210615, token_id: goBlock
	fmt.Println("유저 네임을 입력하세요: ")
	fmt.Scanln(&mySafeNumber.owner)
	fmt.Println("저장할 NFT 토큰 ID를 입력하세요: ")
	fmt.Scanln(&mySafeNumber.token_id)
	fmt.Println("저장할 NFT 토큰의 lock time을 입력하세요: ")
	fmt.Scanln(&mySafeNumber.lock_time)
	fmt.Printf("%v님의 아이템을 저장하는 중입니다...\n", mySafeNumber.owner)

	// Creat a random safe number and assign it to mySafe.
	rand.Seed(time.Now().UnixNano())
	temp_safeNum := rand.Int63n(10)
	mySafe.safeNum = int(temp_safeNum)

	// Calculate knot first.
	radius := mySafeNumber.lock_time
	height := mySafeNumber.token_id
	mySafe.knot = math.Pi * math.Pow(radius, 2) * height

	// Add a lock_time, token_id, and safe.knot to create a safePw.
	mySafeNumber.safePw = radius + height + mySafe.knot

	// Declare a boolean variable to state if item is well secured.
	isSecured := true

	// Print a message: safeNum: xxx, owner: jonghyun item well secured.
	fmt.Printf("Owner:%s\n Safe number:%d\n Safe password: %f\n isSecured: %v\n ", mySafeNumber.owner, mySafe.safeNum, mySafeNumber.safePw, isSecured)
	fmt.Println()
	fmt.Println("Item safely secured.")
	fmt.Println()

	// Check item storage state.
	var user_safe_pw float64
	fmt.Println("NFT를 회수하기 위해 금고 비밀 번호를 입력하세요: ")
	fmt.Scanln(&user_safe_pw)
	stateCheck := mySafeNumber.stateCheck(user_safe_pw)

	if stateCheck == true {
		fmt.Println("NFT가 회수되었습니다.")
	} else {
		log.Fatal("비밀번호가 다릅니다. 프로그램을 종료합니다.")
	}

}

type safe struct {
	isSecured bool
	knot      float64
	dial      float64
	safeNum   int
}

type safePassword struct {
	owner     string
	lock_time float64 // becomes radius of knot's volume formular
	token_id  float64 // becomes height of knot's volume formular
	safePw    float64 // safePw = lock_time + token_id + safe.knot
}

// Create a input & output method

// Create a stateCheck method
func (s_pw *safePassword) stateCheck(pwCheck float64) bool {
	if pwCheck == s_pw.safePw {
		return true
	}
	return false
}
