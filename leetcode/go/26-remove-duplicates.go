package hackerrank

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 0
	for right := 1; right < len(nums); right++ {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
	}
	return left + 1
}