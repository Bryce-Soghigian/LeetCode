class Solution:
    def bstToGst(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        self.curval = 0
        self.traverse(root)
        
        return root
    
    def traverse(self, root):
        if root==None:
            return 
        
        self.traverse(root.right)
        self.curval += root.val
        root.val = self.curval
        self.traverse(root.left)
        