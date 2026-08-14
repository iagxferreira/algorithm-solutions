package leetcode

func searchInsert(nums []int, target int) int {
	i := 0
	for _, v := range nums {
		if v >= target {
			return i
		}
		i++
	}
	return i
}
