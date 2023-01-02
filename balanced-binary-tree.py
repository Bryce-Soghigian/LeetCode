# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    no_bueno = True
    
    def dfs(self, tree):
        if not tree:
            return 0
        
        left_subtree = self.dfs(tree.left)
        right_subtree = self.dfs(tree.right)
        
        if abs(left_subtree-right_subtree) > 1:
            self.no_bueno = False
        
        return max(left_subtree, right_subtree) + 1
    def isBalanced(self, root: Optional[TreeNode]) -> bool:
        
        self.dfs(root)
        return self.no_bueno