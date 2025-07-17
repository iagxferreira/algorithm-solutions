/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	tree := buildTree(nums, 0, len(nums)-1)
	return tree
}

func buildTree(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = buildTree(nums, left, mid-1)
	root.Right = buildTree(nums, mid+1, right)
	return root
}
