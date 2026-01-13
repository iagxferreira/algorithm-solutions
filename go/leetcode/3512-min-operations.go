package leetcode

func minOperations(nums []int, k int) int {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	return sum % k
}
