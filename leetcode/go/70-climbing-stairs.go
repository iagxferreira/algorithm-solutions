func climbStairs(n int) int {
	memo := make(map[int]int)
	memo[1] = 1
	memo[2] = 2

	if value, ok := memo[n]; ok {
		return value
	}

	result := climbStairs(n-1) + climbStairs(n-2)
	return result
}
