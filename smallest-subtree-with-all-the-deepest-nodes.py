# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right


class NodeStorage:
    def __init__(self, height, parentNode):
        self.height = height
        self.parentNode = parentNode
class Solution:
    def subtreeWithAllDeepest(self, root: TreeNode) -> TreeNode:
        def dfs(node, depth=0):
            # (5) Recursion Breack:
            if node is None:
                return None, depth - 1

            # (2) Recursion Step:
            llca, ldep = dfs(node.left, depth + 1)
            rlca, rdep = dfs(node.right, depth + 1)

            # (3) Balacnced subtree -> return yourself:
            if rdep == ldep:
                return node, ldep

            # (4) Unbalacnced subtree -> return deeper side:
            else:
                return (llca, ldep) if ldep > rdep else (rlca, rdep)

        # (1) Initiate post order Traversal:
        return dfs(root)[0]
        
        max_so_far = (0, root)
        
        queue = deque()
        queue.append((None,1,root))
        curr_level = 0
        while queue:
            curr_level += 1
            size = len(queue)
            for i in range(size):
                parent,height,curr = queue.popleft()
                
                if curr.left == None and curr.right == None and height > max_so_far[0]:
                    max_so_far = (height,parent)
                
                if curr.left:
                    queue.append((curr,height+1,curr.left))
                if curr.right:
                    queue.append((curr,height+1,curr.right))
        
        
        return max_so_far[1]