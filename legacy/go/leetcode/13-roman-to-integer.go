package leetcode

func romanToInt(s string) int {
	var total = 0
	hashmap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	for index := range s {
		if index < len(s)-1 && hashmap[s[index]] < hashmap[s[index+1]] {
			total -= hashmap[s[index]]
		} else {
			total += hashmap[s[index]]
		}
	}
	return total
}
