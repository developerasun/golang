package main

import "crypto/sha256"

func main() {
	sha256.New().BlockSize()
	sha256.New().Reset()
}

type safe struct {
	storage float64
	knot    float64
	dial    float64
	safeNum int
}

type safeNumber struct {
	owner     string
	lock_time float64
	token_id  string
	safeNum   string // safeNum = lock_time + token_id + safe.knot
}

type Account struct {
	owner string
	ID    string
}

type Keeping interface {
	// Declare methods for PoSN interface.
	SafeCreate()
	SafeNum_init()
	Open()
	Close()
	Lock()
	Unlock()
	Deposit()
	Withdraw()
	Schedule()
}

// Create several safes randomly for decentralization.
func safeCreate() {

}

// Initialize a safe number and assign it to a safe.
func (s *safeNumber) safeNum_init() float64 {
	return 0
}

// Open a safe
func Open(key string) {

}

// Close a safe
func Close() {

}

// Lock a safe
func Lock() {

}

// Unlock a safe
func Unlock() {

}

// Deposit a NFT
func Deposit() {

}

// Withdraw a NFT
func Withdraw() {

}

// Set a date when to get NFT back to user's inventory.
func Schedule() {

}

// Assign Kyuri & Junseok to below methods
// Deposit, Withdraw

// Assign Taewoong & Jonghyun to below methods
// safeCreate, safeNum_init

// Purpose of this demo program
// 1. User enters NFT token ID
// 2. Create a random safe and assign it to the user
// 3. Safe stores the NFT
// 4. Print safe number and message: item safely secured.
