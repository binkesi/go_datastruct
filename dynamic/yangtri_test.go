package dynamic

import "testing"

func TestYangTri(t *testing.T) {
	triangle := []int{5, 7, 8, 2, 3, 4, 4, 9, 6, 1, 2, 7, 9, 4, 5}
	minPath := YangTri(triangle)
	if minPath != 20 {
		t.Errorf("Min path length is %d", minPath)
	}
}
