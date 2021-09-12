package dynamic

import "fmt"

func PutStone(weight []int, col int) int {
	row := len(weight)
	state := make([][]int, row)
	for i := 0; i < row; i++ {
		state[i][0] = 0
	}
	for j := 0; j <= col; j++ {
		if j == weight[0] {
			state[0][j] = 1
		} else {
			state[0][j] = 0
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j <= col; j++ {
			if state[i-1][j] == 1 || (j-weight[i] > 0 && state[i-1][j-weight[i]] == 1) {
				state[i][j] = 1
			} else {
				state[i][j] = 0
			}
			fmt.Printf("%d row %d col value is %d", row, col, state[i][j])
		}
	}
	for k := col; k >= 0; k-- {
		if state[row-1][k] == 1 {
			fmt.Printf("max value is %d", k)
			return k
		}
	}
	return 0
}
