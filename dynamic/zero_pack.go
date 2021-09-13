package dynamic

import "fmt"

func PutStone(weight []int, col int) int {
	row := len(weight)
	state := [][]int{}
	// init two-d array
	for r := 0; r < row; r++ {
		state = append(state, make([]int, col+1))
	}
	if weight[0] <= col {
		state[0][weight[0]] = 1
	}
	for i := 1; i < row; i++ {
		for j := 1; j <= col; j++ {
			if weight[i] <= col {
				state[i][weight[i]] = 1
			}
			if state[i-1][j] == 1 || (j-weight[i] > 0 && state[i-1][j-weight[i]] == 1) {
				state[i][j] = 1
			} else {
				state[i][j] = 0
			}
		}
	}
	for i := 0; i < row; i++ {
		fmt.Printf("%v\n", state[i])
	}
	for k := col; k >= 0; k-- {
		if state[row-1][k] == 1 {
			fmt.Printf("max value is %d\n", k)
			return k
		}
	}
	return 0
}
