package leetcode

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := ""
	n := len(s)
	diff := 2 * (numRows - 1)
	diagonalDiff := diff
	var secondIndex, index int
	for i := range numRows {
		index = i
		for index < n {
			rows += string(s[index])
			if i != 0 && i != numRows-1 {
				diagonalDiff = diff - 2*i
				secondIndex = index + diagonalDiff
				if secondIndex < n {
					rows += string(s[secondIndex])
				}
			}
			index += diff
		}
	}
	return rows
}
