package main

import (
	"fmt"
	"hash/crc32"
)

// In programming, hash is used in various ways such as looking up data and deteching unwanted changes
// Go provides two ways of hashing : 1) cryptographic 2) non-cryptographic

func main() {
	// Create a hasher: hash function turn a set of data into a smaller fixed size
	// crc32 function is used to compare two files. If the return values of the two files are the same,
	// mostly likely they are the same file. If not, they are 100% guranteed not the same
	file1 := crc32.NewIEEE()
	file2 := crc32.NewIEEE()

	file1.Write([]byte("crc32 - file 1 go function"))
	file2.Write([]byte("crc32 - file 2 go function"))

	// uint : unsigned integer, meaning > 0
	if file1.Sum32() != file2.Sum32() {
		fmt.Println("They are not the same file!")
		return
	} else {
		fmt.Println(true)
		return
	}

}
