func numTrees(n int) int {
	memo := make([]int, n+1)
	memo[0], memo[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			memo[i] += memo[j-1] * memo[i-j]
		}
	}
	return memo[n]
}
