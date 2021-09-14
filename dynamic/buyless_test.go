package dynamic

import "testing"

func TestBuyLess(t *testing.T) {
	merch := [10]int{59, 124, 89, 182, 67, 90, 174, 46, 98, 35}
	//lessMerch := [5]int{59, 124, 89, 182, 67}
	lessVal := 200
	minValue := BuyLess(merch[:], lessVal)
	if minValue != lessVal {
		t.Errorf("Min buy value %d\n", minValue)
	}
}
