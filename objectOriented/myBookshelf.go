package main

import "fmt"

func main() {
	var myBookshelf Bookshelf
	myBookshelf.init()
	myBookshelf.count()
}

// Create an object
type Bookshelf struct {
	horror  int
	romance int
	comedy  int
}

//  Create an method 1 : init()
func (b *Bookshelf) init() {
	b.horror = 10
	b.romance = 7
	b.comedy = 5
}

//  Create an method 2 : count()
func (b *Bookshelf) count() {
	// the number of book categories in bookshelf : 3
	mySlice := []int{}
	mySlice = append(mySlice, b.comedy, b.horror, b.romance)
	total := 0
	for i := 0; i < len(mySlice); i++ {
		total += mySlice[i]
	}
	fmt.Println(total)
}