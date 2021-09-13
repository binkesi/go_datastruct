package dynamic

import "fmt"

func OnePack(Weight []int, col int) int {
	row := len(Weight)
	state := make([]int, col+1)
	if Weight[0] <= col {
		state[Weight[0]] = 1
	}
	for i := 1; i < row; i++ {
		for j := 0; j <= col; j++ {
			if state[j] == 1 && j+Weight[i] <= col {
				state[j+Weight[i]] = 1
			}
			if Weight[i] <= col {
				state[Weight[i]] = 1
			}
		}
	}
	fmt.Printf("Package state result: %v", state)
	for k := col; k >= 0; k-- {
		if state[k] == 1 {
			return k
		}
	}
	return 0
}
