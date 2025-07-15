import "unicode"

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}

	vowels := "aeiou"
	hasVowel, hasConsonant := false, false

	for _, char := range word {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return false
		}

		if unicode.IsLetter(char) {
			if strings.ContainsRune(vowels, unicode.ToLower(char)) {
				hasVowel = true
			} else {
				hasConsonant = true
			}
		}
	}

	return hasVowel && hasConsonant
}
