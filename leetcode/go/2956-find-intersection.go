package leetcode

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	response := make([]int, 2)
	hashmap1 := make(map[int]int)
	hashmap2 := make(map[int]int)
	for _, value := range nums1 {
		hashmap1[value]++
	}
	for _, value := range nums2 {
		hashmap2[value]++
	}
	for key, value := range hashmap1 {
		if hashmap2[key] != 0 {
			response[0] += value
			response[1] += hashmap2[key]
		}
	}
	return response
}
