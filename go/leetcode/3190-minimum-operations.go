package leetcode

func minimumOperations(nums []int) int {
	answer := 0
	for _, num := range nums {
		if num%3 != 0 {
			answer++
		}
	}
	return answer
}
