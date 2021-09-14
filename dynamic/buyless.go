package dynamic

import (
	"fmt"
)

func BuyLess(merch []int, lessVal int) int {
	row := len(merch)
	col := int(2 * lessVal)
	state := make([]int, col+1)
	state[0] = 1
	if merch[0] <= col {
		state[merch[0]] = 1
	}
	for i := 1; i < row; i++ {
		for j := col; j >= 0; j-- {
			if merch[i] <= j && state[j-merch[i]] == 1 {
				state[j] = 1
			}
		}
	}
	for k := lessVal; k <= col; k++ {
		if state[k] == 1 {
			minValue := k
			// get buy list
			buyList := []int{}
			for i := row - 1; i >= 0; i-- {
				if k-merch[i] >= 0 && state[k-merch[i]] == 1 {
					buyList = append(buyList, merch[i])
					k -= merch[i]
				}
			}
			// reverse buy list.
			for i, j := 0, len(buyList)-1; i < j; i, j = i+1, j-1 {
				buyList[i], buyList[j] = buyList[j], buyList[i]
			}
			fmt.Printf("Should buy these merchandise %v\n", buyList)
			return minValue
		}
	}
	return 0
}
