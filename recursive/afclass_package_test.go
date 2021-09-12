package recursive

import (
	"fmt"
	"testing"
)

func TestAfPackage(t *testing.T) {
	Wlimit := 20
	Wlist := [10]int{4, 7, 2, 6, 10, 1, 5, 2, 12, 8}
	Vlist := [10]int{13, 18, 8, 16, 24, 5, 14, 9, 28, 22}
	var Rlist []int
	AfPackage(Wlimit, 0, 0, 0, Wlist[:], Vlist[:], Rlist)
	fmt.Printf("Pick list is %v\n", Plist)
	if maxV != 65 {
		t.Errorf("Max value %d is less than 50", maxV)
	}
}
