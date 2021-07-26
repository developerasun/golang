package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare1(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(81, square(9), "sqaure(9) should be 81")

}

func TestSquare2(t *testing.T) {
	rst := square(3)
	if rst != 9 {
		t.Errorf("Square(3) should be 9 but square(3) returns %d", rst)
	}
}
