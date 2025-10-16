package leetcode

import "math"

func reverse(x int) int {
	signal := 1
	if x < 0 {
		signal = -1
		x = -x
	}

	result := 0
	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}

	result *= signal
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}

	return result
}
