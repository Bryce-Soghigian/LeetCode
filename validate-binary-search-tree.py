# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isValidBST(self, root, minimum=float("inf"),maximum=float("-inf")) -> bool:
        """
        Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
        """
        def traverse(tree,maximum,minimum):
            if tree == None:
                return True
            elif((maximum != None and tree.val >= maximum) or (minimum != None and tree.val <= minimum)):
                return False
            else:
                return traverse(tree.left, tree.val, minimum) and traverse(tree.right, maximum, tree.val)
        return traverse(root, None, None)
                
            
        