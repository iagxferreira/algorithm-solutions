/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverse(node *ListNode) *ListNode {
	var aux *ListNode
	for node != nil {
		temp := node.Next
		node.Next = aux
		aux = node
		node = temp
	}
	return aux
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sumList, auxList *ListNode
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if carry != 0 {
			sum += carry
		}

		carry = sum / 10

		node := &ListNode{Val: sum % 10, Next: nil}
		if sumList == nil {
			sumList = node
			auxList = node
		} else {
			auxList.Next = node
			auxList = auxList.Next
		}
	}

	return sumList
}
