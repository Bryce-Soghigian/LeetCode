"""
# Definition for a Node.
class Node:
    def __init__(self, val: int = 0, left: 'Node' = None, right: 'Node' = None, next: 'Node' = None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next
"""

class Solution:
    def connect(self, root: 'Optional[Node]') -> 'Optional[Node]':
        
        """
        BFS with two pointers
        
        
        
        """
        
        if not root:
            return root
        
        queue = deque([root])
        
        
        while queue:
            
            size = len(queue)
            prev = None
            for i in range(size):
                curr = queue.popleft()
                if prev:
                    prev.next = curr
                if curr.left:
                    queue.append(curr.left)
                
                if curr.right:
                    queue.append(curr.right)
                prev = curr
            prev = None
        return root