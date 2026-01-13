package leetcode

var memo = []int{0, 1, 1}

func fib(n int) int {
	if len(memo) > n || n == 0 {
		return memo[n]
	}
	response := fib(n-1) + fib(n-2)
	memo = append(memo, response)
	return response
}
