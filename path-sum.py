# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def hasPathSum(self, root: Optional[TreeNode], targetSum: int) -> bool:
        if not root:
            return False
        queue = collections.deque([(root,0)])

        while queue:
            curr, sumSoFar = queue.popleft()
            sumSoFar += curr.val
            if not curr.left and not curr.right and sumSoFar == targetSum:
                return True
            if curr.left: queue.append((curr.left, sumSoFar))
            if curr.right: queue.append((curr.right, sumSoFar))




