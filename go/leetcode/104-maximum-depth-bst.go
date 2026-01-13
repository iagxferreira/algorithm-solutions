package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	maxLevel := 0
	inOrderTraversal(root, 1, &maxLevel)
	return maxLevel
}

func inOrderTraversal(root *TreeNode, level int, maxLevel *int) {
	if root == nil {
		return
	}

	if level > *maxLevel {
		*maxLevel = level
	}
	inOrderTraversal(root.Left, level+1, maxLevel)
	inOrderTraversal(root.Right, level+1, maxLevel)
}
