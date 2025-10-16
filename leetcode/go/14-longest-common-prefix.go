package leetcode

import "sort"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	sort.Strings(strs)
	first, last := strs[0], strs[len(strs)-1]
	result := ""
	for i := 0; i < len(first); i++ {
		if !(i < len(last) && first[i] == last[i]) {
			break
		}
		result += string(first[i])
	}
	return result
}
