from typing import Optional


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def rotateRight(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        if not head:
            return None

        last = head
        length = 1

        while last.next:
            last = last.next
            length += 1

        k = k % length

        last.next = head

        temp = head
        for _ in range(length - k - 1):
            temp = temp.next

        answer = temp.next
        temp.next = None
        return answer
