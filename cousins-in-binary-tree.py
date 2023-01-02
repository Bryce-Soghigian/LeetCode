# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isCousins(self, root: Optional[TreeNode], x: int, y: int) -> bool:
        """
                    1
                 
                 2      3
                 
            4
                    
                
        1. BFS Doing a level order keeping track of the parent. 
        
        Each level we traverse we need some context.
        
        seen_x, seen_y = None, None
        
        When we come across a node we want to store its parent
        if after the level order traversal we have seen both and they dont have the 
        same parent we return True else false
        
        if we never see the nodes we return False because they are in a tree in some other forest.
        
        """
        
        queue = deque([(root, None)])
        while queue:
            size = len(queue)
            seen_x = None
            seen_y = None
            
            for _ in range(size):
                curr, parent = queue.popleft()
                if curr is not None:
                    if curr.val == x:
                        seen_x = parent
                    if curr.val == y:
                        seen_y = parent
                    
                if curr.left:
                    queue.append((curr.left, curr))
                if curr.right:
                    queue.append((curr.right, curr))

            if seen_x != None and seen_y != None:
                if seen_x == seen_y:
                    return False
                return True
        return False