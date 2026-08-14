package leetcode

func max(a, b int) int {

	if a > b {
		return a
	}

	return b
}
func longestValidParentheses(s string) int {
	stack := []int{-1}
	maxLength := 0

	for index, char := range s {
		if char == '(' {
			stack = append(stack, index)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, index)
			} else {
				maxLength = max(maxLength, index-stack[len(stack)-1])
			}
		}
	}
	return maxLength
}