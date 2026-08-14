package leetcode

func decodeMessage(key string, message string) string {
	rainbowTable := make(map[rune]rune)
	alphabet := 'a'

	for _, char := range key {
		if char != ' ' && rainbowTable[char] == 0 {
			rainbowTable[char] = alphabet
			alphabet++
		}
	}

	decoded := make([]rune, len(message))
	for i, char := range message {
		if char == ' ' {
			decoded[i] = ' '
		} else {
			decoded[i] = rainbowTable[char]
		}
	}
	return string(decoded)
}
