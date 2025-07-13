func binarySearch(items []int, low, high, target int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if items[mid] == target {
		return mid
	} else if target > items[mid] {
		return binarySearch(items, mid+1, high, target)
	}
	return binarySearch(items, low, mid-1, target)
}

func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}
