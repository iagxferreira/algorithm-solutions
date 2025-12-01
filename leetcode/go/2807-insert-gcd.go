package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	node1 := head
	node2 := head.Next

	for node2 != nil {
		gcd := calculateGCD(node1.Val, node2.Val)
		gcdNode := ListNode{Val: gcd, Next: nil}

		node1.Next = &gcdNode
		gcdNode.Next = node2

		node1 = node2
		node2 = node2.Next
	}

	return head
}

func calculateGCD(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}
