package leetcode

func firstUniqChar(s string) int {
	unique := make(map[rune]int)
	for _, caractere := range s {
		unique[caractere]++
	}

	for index, caractere := range s {
		if unique[caractere] == 1 {
			return index
		}
	}
	return -1
}
