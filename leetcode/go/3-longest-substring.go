func lengthOfLongestSubstring(s string) int {
	right := len(s)
	max := 0
	lastIndex := make([]int, 128)
	start := 0
	for left := 0; left < right; left++ {
		current := s[left]
		if lastIndex[current] > start {
			start = lastIndex[current]
		}
		if left-start+1 > max {
			max = left - start + 1
		}
		lastIndex[current] = left + 1
	}
	return max
}
