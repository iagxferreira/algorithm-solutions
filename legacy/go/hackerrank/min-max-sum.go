package hackerrank

import (
	"fmt"
	"slices"
)

func miniMaxSum(arr []int32) {
    slices.Sort(arr)
    min, max, sum := arr[0], arr[0], int64(0)

    for _, num := range arr {
        sum, min, max = sum+int64(num), minInt(min, num), maxInt(max, num)
    }

    fmt.Println(sum-int64(max), sum-int64(min))
}

func minInt(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
