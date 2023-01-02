class Solution:
    def flatten(self, root: TreeNode) -> None:
        """
        Do not return anything, modify root in-place instead.
        """
        curr = root
        while curr:
            if curr.left:
                dummy = curr.left
                while dummy.right:
                    dummy = dummy.right
                dummy.right, curr.right, curr.left = curr.right, curr.left, None
            curr = curr.right