# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def levelOrderBottom(self, root: TreeNode) -> List[List[int]]:
        if root == None:
            return []
        """
        Perform a bfs where we prepend all of our nodes to the array
        
        """
        results = []
        queue = []
        queue.insert(0,root)
        while len(queue) != 0:
            level_size = len(queue)
            level_nodes = []
            for i in range(level_size):
                current = queue.pop()
                level_nodes.append(current.val)
                if current.left != None:
                    queue.insert(0,current.left)
                if current.right != None:
                    queue.insert(0, current.right)
            results.insert(0,level_nodes)
        return results