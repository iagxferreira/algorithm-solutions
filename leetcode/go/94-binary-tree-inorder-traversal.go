package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inOrder(node *TreeNode, arr []int) []int {
	if node == nil {
		return arr
	}
	arr = inOrder(node.Left, arr)
	arr = append(arr, node.Val)
	arr = inOrder(node.Right, arr)
	return arr
}

func inorderTraversal(root *TreeNode) []int {
	return inOrder(root, []int{})
}
