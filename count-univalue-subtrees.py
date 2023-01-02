class Solution:
    def countUnivalSubtrees(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0
        return self.count_unival_helper(root)[0]
    
    def count_unival_helper(self, root: Optional[TreeNode]) -> (int, int):
        # base case: leaf node
        if root.left is None and root.right is None:
            return 1, root.val
        
        l_count = r_count = 0
        l_val = r_val = root.val
        if root.left:
            l_count, l_val = self.count_unival_helper(root.left)
        if root.right:
            r_count, r_val = self.count_unival_helper(root.right)
        count = l_count + r_count
        val = root.val
        if val == l_val and val == r_val:
            count += 1
        else:
            val = float('inf')    # indicates this subtree is not unival
        return count, val