package leetcode

func maxFreqSum(s string) int {
	hashmap := make(map[rune]int)
	maxvow, maxcon := 0, 0
	for _, value := range s {
		switch value {
		case 'a', 'e', 'i', 'o', 'u':
			hashmap[value]++
			if hashmap[value] > maxvow {
				maxvow = hashmap[value]
			}
		default:
			hashmap[value]++
			if hashmap[value] > maxcon {
				maxcon = hashmap[value]
			}
		}
	}
	return maxcon + maxvow
}
