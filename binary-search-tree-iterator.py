# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class BSTIterator:

    def __init__(self, root: Optional[TreeNode]):
        self.state = []
        self.pt = 0
        self._create_state(root)
        

    def _create_state(self, root):
        if not root:
            return None
        
        if root.left: self._create_state(root.left)
        self.state.append(root.val)
        if root.right: self._create_state(root.right)
            
    def next(self) -> int:
        
            old = self.state[self.pt]
            self.pt += 1
            return old
        

        
    def hasNext(self) -> bool:
        return self.pt <= len(self.state) - 1
        


# Your BSTIterator object will be instantiated and called as such:
# obj = BSTIterator(root)
# param_1 = obj.next()
# param_2 = obj.hasNext()