package main

import "fmt"

func main() {
	// byte belongs to uint8 data type in go. Used to express ASCII.
	var ch1 byte = 65
	var ch2 byte = 0101
	var ch3 byte = 0x41
	var ch4 byte = 'A'

	fmt.Printf("%c %c %c %c\n", ch1, ch2, ch3, ch4)

	// rune belongs to int32 data type in go. Used to express Unicode.
	var cha1 rune = 44032
	var cha2 rune = 0126000
	var cha3 rune = 0xAC00
	var cha4 rune = '가'

	fmt.Printf("%c %c %c %c", cha1, cha2, cha3, cha4)
}
