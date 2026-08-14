package leetcode

func buildArray(nums []int) []int {
	answer := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		answer[i] = nums[nums[i]]
	}
	return answer
}
