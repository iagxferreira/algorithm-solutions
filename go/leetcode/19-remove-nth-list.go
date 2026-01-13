package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	if n == length {
		return head.Next
	}

	node := head
	for i := 0; i < length-n-1; i++ {
		node = node.Next
	}

	node.Next = node.Next.Next
	return head
}
