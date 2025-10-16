package leetcode

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, n := range nums {
		if seen[n] {
			return true
		}
		seen[n] = true
	}
	return false
}
