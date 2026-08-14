package leetcode

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var merge []int
	merge = append(merge, nums1...)
	merge = append(merge, nums2...)
	sort.Slice(merge, func(i, j int) bool {
		return merge[i] < merge[j]
	})

	length := len(merge)
	middle := length / 2

	if length%2 == 0 {
		return float64(merge[middle-1]+merge[middle]) / 2.0
	}

	return float64(merge[middle])
}
