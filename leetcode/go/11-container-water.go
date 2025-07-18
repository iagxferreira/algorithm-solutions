func maxArea(height []int) int {
	length := len(height)
	width := length - 1
	left := 0
	right := width
	maxAreaTemp := 0
	for range length {
		high := min(height[left], height[right])
		maxAreaTemp = max(maxAreaTemp, high*width)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		width--
	}
	return maxAreaTemp
}
