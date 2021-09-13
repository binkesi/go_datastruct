package dynamic

import (
	"testing"
)

func TestOnePack(t *testing.T) {
	weight := [10]int{4, 7, 2, 10, 1, 8, 13, 5, 12, 11}
	Wlimit := 20
	maxW := OnePack(weight[:], Wlimit)
	if maxW != Wlimit {
		t.Errorf("Max weight %d is not 20", maxW)
	}
}
