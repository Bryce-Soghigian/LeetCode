# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def _combine_trees(self, arr1, arr2):
        return arr1 + arr2
    
    def _get_tree_values(self, tree):
        if tree == None:
            return []
        queue = deque([tree])
        output = []
        while queue:
            curr=  queue.popleft()
            output.append(curr.val)
            if curr.left:
                queue.append(curr.left)
            if curr.right:
                queue.append(curr.right)
        return output
    def getAllElements(self, root1: TreeNode, root2: TreeNode) -> List[int]:
        """
        Brute force
        1. Traverse both trees and get all their values
        2. combine two values
        3. sort values
        
        """
        return sorted(self._combine_trees(self._get_tree_values(root1), self._get_tree_values(root2)))        