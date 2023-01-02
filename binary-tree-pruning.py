
class Solution:
    def pruneTree(self, root: TreeNode) -> TreeNode:
        def dfs(tree):
            if tree == None:
                return False
            
            left_subtree = dfs(tree.left)
            right_subtree = dfs(tree.right)
            
            if not left_subtree:
                tree.left = None
            if not right_subtree:
                tree.right = None
            return tree.val or left_subtree or right_subtree
            
        dfs(root)
        
        return root if dfs(root) else None