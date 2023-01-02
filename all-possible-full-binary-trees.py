
class Solution:
    def allPossibleFBT(self, n: int) -> List[Optional[TreeNode]]:
        
        self.cache = {}
        if len(self.cache) == 0:
            self.cache[0] = []
            self.cache[1] = [TreeNode(0)]
        if n % 2 == 0:
            # if even its impossible to generate trees
            return []
        def backtrack(n):
            if n in self.cache:
                return self.cache[n]
            output = []
           
            for i in range(1, n, 2):
                left = backtrack(i)
                right = backtrack(n-1-i)
                for l_node in left:
                    for r_node in right:
                        root = TreeNode(0)
                        root.left = l_node
                        root.right = r_node
                        output.append(root)
            self.cache[n] = output
            return output
        return backtrack(n) 

        