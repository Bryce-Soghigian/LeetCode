class Solution:
    global_max = 0
    
    def dfs(self, tree):
        if not tree:
            return 0
        largest_left = self.dfs(tree.left)
        largest_right = self.dfs(tree.right)
        self.global_max = max(self.global_max, largest_left + largest_right)
        return max(largest_left, largest_right) + 1
    def diameterOfBinaryTree(self, root: TreeNode) -> int:
        self.dfs(root)
        return self.global_max