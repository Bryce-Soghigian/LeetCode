# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def verticalOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        if root is None:
            return []
        """
        
        
        queue = deque()
        
        node, col = queue.popleft()
        
        
        """
        
        
        levels = defaultdict(list)
        
        
        queue = deque([(root, 0)])
        
        while queue:
            
            curr, level = queue.popleft()

            
            
            levels[level].append(curr.val)
            
            if curr.left: queue.append((curr.left, level - 1))
            if curr.right: queue.append((curr.right, level + 1))
                
        
        #1. Map all nodes to a column
        #2 Build the columsn 
        
        levels = [item[1] for item in sorted(levels.items(), key = lambda x: x[0])]
        return levels