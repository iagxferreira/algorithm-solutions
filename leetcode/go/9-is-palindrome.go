package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reverse := 0
	aux := x

	for x > 0 {
		reverse = (reverse * 10) + (x % 10)
		x = x / 10
	}
	return reverse == aux
}