# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from collections import deque
class Solution:
    def maxLevelSum(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        queue = deque()
        queue.append(root)
        level = 0
        max_sum = float('-inf')
        max_level = 0
        while queue:
            curr_sum = 0
            size = len(queue)
            level += 1
            for _ in range(size):
                curr = queue.popleft()
                curr_sum += curr.val
                
                if curr.left:
                    queue.append(curr.left)
                if curr.right:
                    queue.append(curr.right)
            if curr_sum > max_sum:
                max_level = level
                max_sum = curr_sum
        return max_level
                
                