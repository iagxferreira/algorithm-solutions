package leetcode

import "strings"

func isPalindrome(s string) bool {
	flag := true
	for left, right := 0, len(s)-1; left <= right; {
		if isValidChar(string(s[left])) && isValidChar(string(s[right])) {
			if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
				return false
			}
			left++
			right--
		} else if !isValidChar(string(s[left])) {
			left++
		} else if !isValidChar(string(s[right])) {
			right--
		}
	}
	return flag
}

func isValidChar(value string) bool {
	if (value >= "A" && value <= "Z") || (value >= "a" && value <= "z") || (value >= "0" && value <= "9") {
		return true
	}
	return false
}
