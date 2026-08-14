package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	smap := make(map[byte]int, 0)
	for i := range len(s) {
		smap[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		smap[t[i]]--
		if smap[t[i]] == 0 {
			delete(smap, t[i])
		} else if smap[t[i]] < 0 {
			return false
		}
	}

	return len(smap) == 0
}
