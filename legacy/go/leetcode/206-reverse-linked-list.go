package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var node *ListNode
	for head != nil {
		temp := head.Next
		head.Next = node
		node = head
		head = temp
	}

	return node
}
