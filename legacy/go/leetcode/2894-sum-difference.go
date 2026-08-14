package leetcode

func differenceOfSums(n int, m int) int {
	answer := 0
	for i := 1; i <= n; i++ {
		if i%m == 0 {
			answer -= i
		} else {
			answer += i
		}
	}
	return answer
}
