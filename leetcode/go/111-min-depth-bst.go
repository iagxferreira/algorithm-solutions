/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	minLevel := math.MaxInt
	inOrderTraversal(root, 1, &minLevel)
	return minLevel
}

func inOrderTraversal(node *TreeNode, level int, minLevel *int) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil && level < *minLevel {
		*minLevel = level
	}

	inOrderTraversal(node.Left, level+1, minLevel)
	inOrderTraversal(node.Right, level+1, minLevel)
}
