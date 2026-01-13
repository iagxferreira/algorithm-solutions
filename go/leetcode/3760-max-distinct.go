package leetcode

func maxDistinct(s string) int {
	seen := make(map[rune]struct{})

	for _, letter := range s {
		seen[letter] = struct{}{}
	}

	return len(seen)
}
