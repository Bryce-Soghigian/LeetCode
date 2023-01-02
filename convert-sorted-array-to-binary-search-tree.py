# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sortedArrayToBST(self, nums: List[int]) -> TreeNode:
        def dfs(left,right):
            if left > right:
                return None
            
            pointer = (left + right) // 2
            
            if (left + right) % 2:
                pointer += 1
                
            root = TreeNode(nums[pointer])
            
            root.left = dfs(left, pointer - 1)
            root.right = dfs(pointer + 1, right)
            return root
        
        return dfs(0, len(nums)-1)
        