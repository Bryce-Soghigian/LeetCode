# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def verticalTraversal(self, root: TreeNode) -> List[List[int]]:
        res = []
        dictionary = dict()
        
        def function(root, vertical, horizontal):
            if not root:
                return 
            
            if vertical in dictionary:
                dictionary[vertical].append((horizontal, root.val))
            else:
                dictionary[vertical] = [(horizontal, root.val)]
            function(root.left, vertical - 1, horizontal + 1)
            function(root.right, vertical + 1, horizontal + 1)

        function(root, 0, 0)     
        # structure of dictionary = {index:[(level, value)]}
        #print(dictionary)
        for i in sorted(dictionary.keys()): # Sorting the keys(indices)
            temp = [j[1] for j in sorted(dictionary[i])] # Sorting in case of a tie
            res.append(temp)
           
        return res