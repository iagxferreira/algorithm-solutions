func permute(nums []int) [][]int {
	res := [][]int{{}}

	for _, num := range nums {
		var newPerms [][]int
		for _, perm := range res {
			for i := 0; i <= len(perm); i++ {
				newPerm := make([]int, len(perm)+1)
				copy(newPerm, perm[:i])
				newPerm[i] = num
				copy(newPerm[i+1:], perm[i:])
				newPerms = append(newPerms, newPerm)
			}
		}
		res = newPerms
	}

	return res
}
