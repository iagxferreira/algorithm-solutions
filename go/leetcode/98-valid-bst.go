package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	memo := inOrderTraverse(root, []int{})

	for i := 0; i < len(memo)-1; i++ {
		if memo[i] >= memo[i+1] {
			return false
		}
	}
	return true
}

func inOrderTraverse(root *TreeNode, inorder []int) []int {
	if root == nil {
		return inorder
	}

	inorder = inOrderTraverse(root.Left, inorder)
	inorder = append(inorder, root.Val)
	inorder = inOrderTraverse(root.Right, inorder)
	return inorder
}
