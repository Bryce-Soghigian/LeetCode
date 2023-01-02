"""
# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children if children is not None else []
"""

class Solution:
    def __init__(self):
        maxdiameter = 0
    def diameter(self, root):
        """
        :type root: 'Node'
        :rtype: int
        """
        self.maxdiameter = 0
        
        def dfs(root):
            depths = [0, 0] # to deal with empty children
            for child in root.children:
                depths.append(dfs(child))
            self.maxdiameter = max(self.maxdiameter, sum(sorted(depths)[-2:])) # sum of top 2 depth
            return max(depths)+1

        dfs(root)
        return self.maxdiameter 