# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def findMode(self, root: Optional[TreeNode]) -> List[int]:
        tree_freq = defaultdict(int)
        
        max_so_far = 0
        queue = deque([root])
        while queue:
            curr = queue.popleft()
            tree_freq[curr.val] += 1
            max_so_far = max(tree_freq[curr.val], max_so_far)
            if curr.left:
                queue.append(curr.left)
            if curr.right:
                queue.append(curr.right)
                
        return [key for key, value in tree_freq.items() if value == max_so_far]