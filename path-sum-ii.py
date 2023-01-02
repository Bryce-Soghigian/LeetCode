# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def __init__(self):
        self.target_sum = 0
    def pathSum(self, root: TreeNode, targetSum: int) -> List[List[int]]:
        if not root:
            return []
        """
        Dfs keeping track of our current path
        keep a running sum of our current path
        
        
        if we are at a leaf node and our running sum == target we want to append currPath to output
        
        
        1. Dfs
        1.
        1.a Keep running sum of path
        1.b Keep track of our current path
        1bc. CHeck to see if curr path is a target path
        1.c return output
        
        
        """
        self.target_sum = targetSum
        def dfs(tree, currPath=[], runningSum=0,output = []):
            
            if tree is None:
                currPath = []
                return 
            
            runningSum += tree.val
            currPath.append(tree.val)
            if tree.left == None and tree.right == None and runningSum == self.target_sum:
                output.append(currPath[:])
            if tree.left:
                dfs(tree.left,currPath,runningSum, output)
                currPath.pop()
            if tree.right:
                dfs(tree.right,currPath,runningSum, output)
                currPath.pop()
            return output
        return dfs(root)