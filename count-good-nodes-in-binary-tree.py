# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def goodNodes(self, root: TreeNode) -> int:
        
        
        queue = deque([(root, float('-inf'))])
        
        good_bois = 0
        while queue:
            curr, max_so_far = queue.popleft()
            
            if curr.val >= max_so_far:
                good_bois += 1
                max_so_far = curr.val
                
            if curr.left: queue.append((curr.left, max_so_far))
            if curr.right: queue.append((curr.right, max_so_far))
        return good_bois