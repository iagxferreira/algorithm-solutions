from typing import Optional


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        stack = [root.right, root.left]

        while stack:
            left_node = stack.pop()
            right_node = stack.pop()

            if not left_node and not right_node:
                continue

            if not left_node or not right_node:
                return False

            if left_node.val != right_node.val:
                return False

            stack.append(right_node.right)
            stack.append(left_node.left)
            stack.append(right_node.left)
            stack.append(left_node.right)
        return True
