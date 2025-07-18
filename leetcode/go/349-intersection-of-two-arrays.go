func intersection(nums1 []int, nums2 []int) []int {
	seen := make(map[int]bool)
	for _, n := range nums1 {
		seen[n] = true
	}

	res := []int{}
	for _, n := range nums2 {
		if seen[n] {
			res = append(res, n)
			delete(seen, n)
		}
	}
	return res
}
