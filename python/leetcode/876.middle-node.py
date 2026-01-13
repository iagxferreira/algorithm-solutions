# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def middleNode(self, head: Optional[ListNode]) -> Optional[ListNode]:
        left = head
        right = head
        while right is not None and right.next is not None:
            left = left.next
            right = right.next.next
        return left
