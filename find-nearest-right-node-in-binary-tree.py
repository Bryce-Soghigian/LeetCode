# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def findNearestRightNode(self, root: TreeNode, u: TreeNode) -> Optional[TreeNode]:
        
        
        queue = deque([root])
        
        while queue:
            
            level_size = len(queue)
            seen_node = False
            for i in range(level_size):
                curr = queue.popleft()
                if i == level_size - 1 and curr == u:
                    return None
                if seen_node: return curr
                if curr == u:
                    seen_node = True
                if curr.left: queue.append(curr.left)
                if curr.right: queue.append(curr.right)
                    
                