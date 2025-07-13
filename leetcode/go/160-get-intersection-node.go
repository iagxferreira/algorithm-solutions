/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	seen := make(map[*ListNode]bool, 0)
	for n := headA; n != nil; n = n.Next {
		seen[n] = true
	}

	for n := headB; n != nil; n = n.Next {
		if seen[n] {
			return n
		}
	}
	return nil
}
