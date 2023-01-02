# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    summation = 0
    def dfs(self,tree, path):
        if tree is None:
            return
        
        path += str(tree.val)
        if not tree.left and not tree.right:
            self.summation += int(path)
            
        if tree.left: self.dfs(tree.left, path)
        if tree.right: self.dfs(tree.right, path)
            
            
    def sumNumbers(self, root: Optional[TreeNode]) -> int:
        self.dfs(root,"")
        return self.summation
        