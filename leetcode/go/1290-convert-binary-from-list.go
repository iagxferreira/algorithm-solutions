/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	var result int
	for head != nil {
		result <<= 1
		result |= head.Val
		head = head.Next
	}
	return result
}
