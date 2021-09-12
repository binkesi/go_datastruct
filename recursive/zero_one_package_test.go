package recursive

import (
	"testing"
)

func TestZerone(t *testing.T) {
	Ulimit := 40
	things := [10]int{5, 8, 2, 14, 9, 7, 10, 6, 21, 4}
	Zerone(Ulimit, things[:], 0, 0)
	if maxW != 40 {
		t.Errorf("Max package value %d is wrong", maxW)
	}
}
