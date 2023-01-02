# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    output = 0
    def tilted_gamr(self, tree):
        if tree is None:
            return 0
        
        left_subtree = self.tilted_gamr(tree.left)
        right_subtree = self.tilted_gamr(tree.right)
        self.output += abs(left_subtree - right_subtree)
        
        return left_subtree + right_subtree + tree.val
    def findTilt(self, root: Optional[TreeNode]) -> int:
        self.tilted_gamr(root)
        return self.output
        
        
        