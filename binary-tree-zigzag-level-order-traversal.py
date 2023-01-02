# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def zigzagLevelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        if root is None:
            return []
        
        
        queue = deque([root])
        output = []
        direction = 0
        while queue:
            size = len(queue)
            curr_bucket = []
            for _ in range(size):
                curr = queue.popleft()
                curr_bucket.append(curr.val)
                if curr.left: queue.append(curr.left)
                if curr.right: queue.append(curr.right)
            if direction ^ 1 == 0:
                output.append(curr_bucket[::-1])
            else:
                output.append(curr_bucket)
            direction = direction ^ 1
        return output
                
                