# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isSymmetric(self, tree: Optional[TreeNode]) -> bool:
            queue = deque([tree])
            queue.append(tree)

            while queue:
                node_one, node_two = queue.popleft(), queue.popleft()
                if node_one == None and node_two == None: continue
                if node_one == None or node_two == None: return False
                if node_one.val != node_two.val: return False

                queue.append(node_one.left)
                queue.append(node_two.right)
                queue.append(node_one.right)
                queue.append(node_two.left)
            return True
