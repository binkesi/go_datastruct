package recursive

import (
	"fmt"
)

// result is list of col index, length is 8.
var res [8]int

func IsValid(row int, col int) bool {
	for i := row - 1; i >= 0; i -= 1 {
		v := res[i]
		k := row - i
		if v == col || v == col-k || v == col+k {
			return false
		}
	}
	return true
}

func PutQueen(row int) {
	if row == 8 {
		fmt.Printf("%v\n", res)
		return
	}
	for i := 0; i <= 7; i++ {
		if IsValid(row, i) {
			res[row] = i
			PutQueen(row + 1)
		}
	}
}
