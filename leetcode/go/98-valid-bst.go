/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func validateWithMemo(node *TreeNode, min *int, max *int, memo []bool) []bool {
	if node == nil {
		return memo
	}

	isValid := true

	if min != nil && node.Val <= *min {
		isValid = false
	}

	if max != nil && node.Val >= *max {
		isValid = false
	}

	memo = append(memo, isValid)

	memo = validateWithMemo(node.Left, min, &node.Val, memo)
	memo = validateWithMemo(node.Right, &node.Val, max, memo)

	return memo
}

func isValidBST(root *TreeNode) bool {
	memo := make([]bool, 0)
	memo = validateWithMemo(root, nil, nil, memo)

	for _, result := range memo {
		if !result {
			return false
		}
	}
	return true
}
