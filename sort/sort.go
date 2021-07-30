package main

import (
	"fmt"
	"sort"
)

func main() {

	// Create slices : int slice, string slice
	intSlice1 := []int{1, 5, 2, 4, 0}
	intSlice2 := []int{-2, 3, -5, 8, 0}
	stringSlice := []string{"a", "c", "f", "g"}

	// Add data into the slice
	intSlice1 = append(intSlice1, 9)
	intSlice2 = append(intSlice2, -22)
	stringSlice = append(stringSlice, "p")

	// Sort the string and int slice
	sort.Strings(stringSlice)
	sort.Ints(intSlice2)
	isStringSorted := sort.StringsAreSorted(stringSlice)

	// Check the intSlices are sorted
	isInt1Sorted := sort.IntsAreSorted(intSlice1)
	isInt2Sorted := sort.IntsAreSorted(intSlice2)

	// Print the sorting result
	fmt.Println("Are the int slices sorted?: ", isInt1Sorted, isInt2Sorted)
	fmt.Println("Is the string slice sorted?: ", isStringSorted)

}
