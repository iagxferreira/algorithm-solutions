package leetcode

import "math"

func scoreOfString(s string) int {
	total := 0.0
	length := len(s)
	for i := 0; i < length-1; i++ {
		total += math.Abs(float64(s[i]) - float64(s[i+1]))
	}
	return int(total)
}
