package dynamic

import (
	"fmt"
	"math"
)

func YangTri(triangle []int) int {
	level := 2
	triangle = append([]int{0}, triangle...)
	for i := 2; i < len(triangle); level++ {
		i += level
	}
	level -= 1
	pathLen := make([]int, len(triangle))
	pathLen[1] = triangle[1]
	for i := 2; i <= level; i++ {
		for j := i * (i + 1) / 2; j >= (i*(i+1)/2 - i + 1); j-- {
			if j == i*(i+1)/2 {
				pathLen[j] = pathLen[j-i] + triangle[j]
			} else if j == (i*(i+1)/2 - i + 1) {
				pathLen[j] = pathLen[j-i+1] + triangle[j]
			} else {
				pathLen[j] = int(math.Min(float64(pathLen[j-i]), float64(pathLen[j-i+1]))) + triangle[j]
			}
		}
	}
	fmt.Printf("%v\n", pathLen)
	minValue := pathLen[len(pathLen)-1]
	for k := len(pathLen) - 2; k >= len(pathLen)-5; k-- {
		if pathLen[k] < minValue {
			minValue = pathLen[k]
		}
	}
	return minValue
}
