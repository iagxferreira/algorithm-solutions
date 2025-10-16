package leetcode

import "sort"

func topKFrequent(nums []int, k int) []int {
	memo := make(map[int]int)
	for _, value := range nums {
		memo[value] += 1
	}

	count := make([][]int, 0)
	for key, value := range memo {
		count = append(count, []int{key, value})
	}

	sort.Slice(count, func(i int, j int) bool {
		return count[i][1] > count[j][1]
	})

	result := make([]int, 0)
	for i := 0; i < k; i++ {
		result = append(result, count[i][0])
	}

	return result
}
