package leetcode

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	nums1Set := make(map[int]bool)
	nums2Set := make(map[int]bool)

	for _, num := range nums1 {
		nums1Set[num] = true
	}
	for _, num := range nums2 {
		nums2Set[num] = true
	}

	ainb := 0
	for _, num := range nums1 {
		if nums2Set[num] {
			ainb++
		}
	}

	bina := 0
	for _, num := range nums2 {
		if nums1Set[num] {
			bina++
		}
	}

	return []int{ainb, bina}
}
