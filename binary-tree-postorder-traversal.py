# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def postorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        """
        
        
        while stack:
            if not stack:
                
            curr = stack.pop()
            if curr.left:
            stack.append()
        []
            res, stack = [], [root]
    while stack:
        node = stack.pop()
        if node:
            res.append(node.val)
            stack.append(node.left)
            stack.append(node.right)
        
        """
        output = []
        stack = deque()
        stack.append(root)
        while stack:
            curr = stack.pop()
            if curr:
                output.append(curr.val)
                stack.append(curr.left)
                stack.append(curr.right)
        return reversed(output)