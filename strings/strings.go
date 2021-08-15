package main

import (
	"fmt"
	"strings"
)

func main() {
	// Prepare variables that will be used for strings package function
	testing := "testing"
	testingVer := "ver 1.0"
	var mySlice []string

	// strings function #1 : strings.Contains
	isContained := strings.Contains(testing, "ing")
	fmt.Println("Is 'ing' in 'testing'? :", isContained)

	// strings function #2 : strings.Count
	isCount := strings.Count(testing, "t")
	fmt.Println("How many 't' is in 'testing'? :", isCount)

	// strings function #3 : strings.HasPrefix
	hasPrefix := strings.HasPrefix(testing, "tes")
	fmt.Println("Does 'testing' have 'tes' as prefix?: ", hasPrefix)

	// strings function #4 : strings.HasSuffix
	hasSuffix := strings.HasSuffix(testing, "ng")
	fmt.Println("Does 'testing' have 'ng' as suffix?: ", hasSuffix)

	// strings function #5 : strings.Index
	idx := strings.Index(testing, "i")
	fmt.Println("Index number of 'i' is: ", idx)

	// strings function #6 : strings.Join
	join := strings.Join([]string{testing, testingVer}, "_")
	fmt.Println(join)

	// strings function #7 : strings.Repeat
	repeat := strings.Repeat("=", 5)
	mySlice = append(mySlice, join, repeat, "by Jake")
	join = strings.Join(mySlice, " ")
	fmt.Println(join)

	// strings function #8 : strings.Replace
	replace := strings.Replace(testing, "t", "wow", 2)
	fmt.Println(replace)

	// strings function #9 : strings.Split
	split := strings.Split(testing, "t")
	for _, val := range split {
		fmt.Println(val)
	}

	// strings function #10 : strings.ToUpper, ToLower
	upper := strings.ToUpper(testing)
	fmt.Println(upper)
	lower := strings.ToLower(upper)
	fmt.Println(lower)

	// Encoding of Strings <-> Binary : necessary for transmission of data
	// Convert strings into binary data - a slice of bytes
	arr := []byte(testing)
	fmt.Println(arr) // A slice of bytes can be converted into several types of formats, depending on its needs

	// Convert binary data into strings
	// A binary-to-text encoding is encoding of data in plain text.
	// More precisely, it is an encoding of binary data in a sequence of printable characters.
	// when the channel does not allow binary data (such as email or NNTP) or is not 8-bit clean.
	str1 := string(arr)
	str2 := string([]byte{'t', 'e', 's', 't', 'i', 'n', 'g'})
	fmt.Println(str1, str2)
}
