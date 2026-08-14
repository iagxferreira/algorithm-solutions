package leetcode

func removeElement(nums []int, val int) int {
	aux := 0
	size := len(nums)
	for i := 0; i < size; i++ {
		if nums[i] != val {
			nums[i], nums[aux] = nums[aux], nums[i]
			aux++
		}
	}
	return aux
}
