# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from collections import deque
class Solution:
    def largestValues(self, root: TreeNode) -> List[int]:
        
        output = []
        if root == None:
            return output
        
        """
        Loop through tree level by level
        at each level we grab the max 
        
        
        """
        
        
        queue = deque()
        
        queue.append(root)
        
        while queue:
            size = len(queue)
            maximum = float("-inf")
            for i in range(size):
                curr = queue.popleft()
                
                maximum = max(curr.val, maximum)
                
                if curr.left:
                    queue.append(curr.left)
                if curr.right:
                    queue.append(curr.right)
            output.append(maximum)
        return output
                
                