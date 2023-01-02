class Solution:
    def rangeSumBST(self, root: Optional[TreeNode], low: int, high: int) -> int:
        def dfs(root):
            total = 0
            if not root: return 0
            if low <= root.val <= high:
                total += root.val
            total += dfs(root.left) + dfs(root.right)
            return total 
        return dfs(root)