# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def kthSmallest(self, root: TreeNode, k: int) -> int:
        inorder_results = []
        def inorder(tree):
            if tree == None:
                return
            
            if tree.left:
                inorder(tree.left)
            inorder_results.append(tree.val)
            if tree.right:
                inorder(tree.right)
        
        inorder(root)
        return inorder_results[k-1]