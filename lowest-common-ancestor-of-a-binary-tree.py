# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def _find_ancestors(self, root, p, q):
        if not root:
            return [False, False]
        if self.output != None:
            return [True, True]
        left_state = self._find_ancestors(root.left, p, q) # [SeenP,SeenQ]
        right_state = self._find_ancestors(root.right, p, q)
        
        seen_p = False if (left_state[0] == False and right_state[0] == False) else True
        seen_q = False if (right_state[1] == False and left_state[1] == False) else True
        
        if root == p:
            seen_p = True
        if root == q:
            seen_q = True
        if self.output == None and seen_p and seen_q:
            self.output = root
            return [True, True]
        return [seen_p, seen_q]
        
        
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        """
        
        Post order dfs
        
                    
                    
                    
        We want to see if     
        
        
        
        base case (seen_p, seen_q)
        
        Find(p)
        find(q)
        if found_p and foundq in tree:
            return True
            
        
        """
        self.output = None
        self._find_ancestors(root, p, q)
        return self.output