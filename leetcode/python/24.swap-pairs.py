# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        placeholder = ListNode(0, head)
        previous_node, current_node = placeholder, head
        while current_node and current_node.next:
            next_pair_node = current_node.next.next
            second = current_node.next

            second.next  = current_node
            current_node.next = next_pair_node
            previous_node.next = second

            previous_node = current_node
            current_node = next_pair_node
        return placeholder.next
