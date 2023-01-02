# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
        
        if not p and q:
            return False
        if not q and p:
            return False
        if not p and not q:
            return True
        queue = deque([(p, q)])
        
        
        while queue:
            
            tree_one, tree_two = queue.popleft()
            if tree_one.val != tree_two.val:
                return False
            
            
            if (tree_one.left and not tree_two.left) or (not tree_one.left and tree_two.left) or (tree_one.right and not tree_two.right) or (not tree_one.right and tree_two.right):
                return False
            
            if tree_one.left:
                queue.append((tree_one.left, tree_two.left))
            
            if tree_one.right:
                queue.append((tree_one.right, tree_two.right))
        return True