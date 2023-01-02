class Solution:
    def findLeaves(self, root: TreeNode) -> List[List[int]]:
        def assignNodeLevel(node, leafNodes) -> int:
            if node is None:
                return 0
            
            leftLevel = assignNodeLevel(node.left, leafNodes)
            rightLevel = assignNodeLevel(node.right, leafNodes)
        
            level = max(leftLevel, rightLevel) + 1
            leafNodes[level].append(node.val)
            
            return level
                
        leafNodes: defaultdict[int, List[TreeNode]] = defaultdict(lambda: [])
        assignNodeLevel(root, leafNodes)
        
        answer = []
        for level in range(1, max(leafNodes.keys()) + 1):
            answer.append(leafNodes[level])
            
        return answer
            
        
        
        