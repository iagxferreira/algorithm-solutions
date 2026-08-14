package leetcode

func findWordsContaining(words []string, x byte) []int {
	response := make([]int, 0)
	for index, word := range words {
		for _, element := range word {
			if byte(element) == x {
				response = append(response, index)
				break
			}
		}
	}
	return response
}
