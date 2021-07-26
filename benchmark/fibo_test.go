package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonnaci1(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, fibonacci1(-1), "fibonnaci1(-1) should be 0")
	assert.Equal(0, fibonacci1(0), "fibonnaci1(0) should be 0")
	assert.Equal(1, fibonacci1(1), "fibonnaci1(1) should be 1")
	assert.Equal(2, fibonacci1(3), "fibonnaci1(2) should be 2")
	assert.Equal(233, fibonacci1(13), "fibonnaci1(13) should be 233")

}

func TestFibonacci2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(0, fibonacci1(-1), "fibonnaci1(-1) should be 0")
	assert.Equal(0, fibonacci1(0), "fibonnaci1(0) should be 0")
	assert.Equal(1, fibonacci1(1), "fibonnaci1(1) should be 1")
	assert.Equal(2, fibonacci1(3), "fibonnaci1(2) should be 2")
	assert.Equal(233, fibonacci1(13), "fibonnaci1(13) should be 233")
}

func BenchmarkFibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci2(20)
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci2(20)
	}
}
