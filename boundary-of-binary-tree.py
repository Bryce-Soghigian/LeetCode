# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    
    def isLeaf(self, root: 'TreeNode') -> bool:
        return True if not root.left and not root.right else False
    
    def getRightBoundary(self, root: 'TreeNode') -> List[int]:
        if not root:
            return []
        temp = root
        res = []
        while temp:
            if not self.isLeaf(temp):
                res.append(temp.val)
            if not temp.right:
                if temp.left:
                    temp = temp.left
                else:
                    break
            else:
                temp = temp.right
        return res
    
    def getLeftBoundary(self, root: 'TreeNode') -> List[int]:
        if not root:
            return []
        temp = root
        res = []
        while temp:
            if not self.isLeaf(temp):
                res.append(temp.val)
            if not temp.left:
                if temp.right:
                    temp = temp.right
                else:
                    break
            else:
                temp = temp.left
        return res
    def dfs(self, root: 'TreeNode', leaves: List[int]) -> None:
        if not root:
            return
        if self.isLeaf(root):
            leaves.append(root.val)
        self.dfs(root.left, leaves)
        self.dfs(root.right, leaves)
        
    
    def boundaryOfBinaryTree(self, root: Optional[TreeNode]) -> List[int]:
        leftBoundary = self.getLeftBoundary(root.left)
        rightBoundary = self.getRightBoundary(root.right)
        leaves = []
        res = [root.val]
        self.dfs(root.left, leaves)
        self.dfs(root.right, leaves)
        rightBoundary.reverse()
        return itertools.chain(res, leftBoundary, leaves, rightBoundary)