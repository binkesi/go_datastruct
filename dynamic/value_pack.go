package dynamic

import (
	"fmt"
	"math"
)

func ValuePack(weight []int, value []int, wlimit int) int {
	row := len(weight)
	state := make([]int, wlimit+1)
	if weight[0] <= wlimit {
		for k := weight[0]; k <= wlimit; k++ {
			state[k] = value[0]
		}
	}
	for i := 1; i < row; i++ {
		for j := wlimit; j >= 0; j-- {
			if j > 0 && state[j] < state[j-1] {
				state[j] = state[j-1]
			}
			if weight[i] <= j {
				state[j] = int(math.Max(float64(state[j]), float64(state[j-weight[i]]+value[i])))
			}
		}
		if weight[i] <= wlimit {
			state[weight[i]] = int(math.Max(float64(value[i]), float64(state[weight[i]])))
		}
		fmt.Printf("%v\n", state)
	}
	return state[wlimit]
}
