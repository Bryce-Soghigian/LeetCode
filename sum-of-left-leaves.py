# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    leaf_sum = 0
    def bfs(self, tree):
        
        
        queue = deque([(tree, False)]) # tree, is_left
        
        while queue:
            curr, is_left_child = queue.popleft()
            if is_left_child and (not curr.left and not curr.right):
                self.leaf_sum += curr.val
            if curr.left: queue.append((curr.left, True))
            if curr.right: queue.append((curr.right, False))
    def sumOfLeftLeaves(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        self.bfs(root)
        return self.leaf_sum