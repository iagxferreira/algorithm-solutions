package hackerrank

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	result := 0
	sign := 1
	index := 0
	n := len(s)

	if n == 0 {
		return 0
	}

	if index < n && s[index] == '-' || s[index] == '+' {
		if s[index] == '-' {
			sign = -1
		}
		index++
	}

	for index < n && s[index] >= '0' && s[index] <= '9' {
		digit := int(s[index] - '0')

		if result > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		result = result*10 + digit
		index++
	}
	return result * sign
}