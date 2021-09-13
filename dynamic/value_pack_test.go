package dynamic

import "testing"

func TestValuePack(t *testing.T) {
	Weight := [10]int{4, 7, 2, 10, 1, 8, 13, 5, 12, 11}
	Value := [10]int{18, 30, 8, 42, 3, 34, 51, 21, 48, 44}
	wlimit := 20
	maxValue := ValuePack(Weight[:], Value[:], wlimit)
	if maxValue != 85 {
		t.Errorf("Max value is %d", maxValue)
	}
}
