class Solution:
    def sumRootToLeaf(self, root: TreeNode) -> int:
        
        def helper(node, num):
            if not node:
                return 0
            num = (num * 2) + node.val
            if not node.left and not node.right:
                return num
            return helper(node.left, num) + helper(node.right, num)
        
        return helper(root, 0)