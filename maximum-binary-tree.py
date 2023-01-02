# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def recurse(self, nums):
        if len(nums) == 0:
            return None
        max_index = None
        max_val = float('-inf')
        for i in range(len(nums)):
            if nums[i] > max_val:
                max_index = i
                max_val = nums[i]
        left_subtree = self.recurse(nums[:max_index])
        right_subtree = self.recurse(nums[max_index + 1:])
        return TreeNode(val=max_val,left= left_subtree,right= right_subtree)
    
    def constructMaximumBinaryTree(self, nums):
        return self.recurse(nums)