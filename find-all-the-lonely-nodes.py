# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:

            
    def getLonelyNodes(self, root: Optional[TreeNode]) -> List[int]:

        if not root:
            return []
        output = []
        queue = deque([(root, 2)])
        while queue:
            curr, siblings = queue.popleft()
            
            if siblings == 1:
                output.append(curr.val)
            new_siblings = 0
            if curr.left: 
                new_siblings += 1
            if curr.right:
                new_siblings += 1
            
            if curr.left:
                queue.append((curr.left, new_siblings))
            
            if curr.right:
                queue.append((curr.right, new_siblings))
        return output