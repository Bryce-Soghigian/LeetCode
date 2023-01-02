"""
# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""

class Solution:
    def levelOrder(self, root: 'Node') -> List[List[int]]:
        if not root:
            return []
        output = []
        queue = deque()
        
        queue.append(root)
        while queue:
            current_level = []
            size = len(queue)
            for i in range(len(queue)):
                curr = queue.popleft()
                
                current_level.append(curr.val)
                
                for child in curr.children:
                    queue.append(child)
            output.append(current_level)
        return output