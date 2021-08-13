package main

import (
	"runtime"
)

type data struct {
	i, j int
}

func main() {
	var N = 40000000
	var structure []data
	for i := 0; i < N; i++ {
		value := int(i)
		structure = append(structure, data{value, value})
	}
	// GC runs a garbage collection and blocks the caller until the garbage collection is complete.
	// It may also block the entire program.
	runtime.GC()

	// Below code is used to prevent Go garbage collector from garbage-collecting the data struct too early
	// even before it is referenced to.
	_ = structure[0]
}
