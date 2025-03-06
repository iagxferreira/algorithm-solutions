package hackerrank

func staircase(n int32) {
	output := ""

	for i := int32(1); i <= n; i++ {
		for s := int32(n - 1); s >= i; s-- {
			output += " "
		}
		for h := int32(1); h <= i; h++ {
			output += "#"
		}
		output += "\n"
	}

	fmt.Print(output)
}