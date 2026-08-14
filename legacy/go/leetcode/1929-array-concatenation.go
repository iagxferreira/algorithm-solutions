package leetcode

func getConcatenation(nums []int) []int {
	length := len(nums)
	answer := make([]int, length*2)
	for i := 0; i < length; i++ {
		answer[i], answer[i+length] = nums[i], nums[i]
	}
	return answer
}
