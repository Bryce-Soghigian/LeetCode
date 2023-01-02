class TreeAncestor:

    def __init__(self, n: int, parent: List[int]):
        self.parent = parent
        
        # Get childrens array
        self.children = [[] for _ in range(len(parent))]
        for i in range(1, len(parent)):
            self.children[parent[i]].append(i)
            
        # Now do backtracking via childrens array to get
        # every single link from head to leaf in the array
        # In addition, we also save which index we save the leaf in!
        self.headToLeafs = []
        self.idxToLookupHeadToLeaf = [0 for _ in range(len(parent))]
        
        def dfsFillUpHeadToLeaf(idxNode, curArray):
            curArray.append(idxNode)
            
            if self.children[idxNode]:
                for childrenIdx in self.children[idxNode]:
                    idxLookup = dfsFillUpHeadToLeaf(childrenIdx, curArray)
            else:
                idxLookup = len(self.headToLeafs)
                self.headToLeafs.append(curArray[:])
            
            self.idxToLookupHeadToLeaf[idxNode] = (idxLookup, len(curArray) - 1)
            curArray.pop()
            return idxLookup
            
        dfsFillUpHeadToLeaf(0, [])
        

    def getKthAncestor(self, node: int, k: int) -> int:
        headToLeafChain = self.headToLeafs[self.idxToLookupHeadToLeaf[node][0]]
        idxInHeadToLeafChain = self.idxToLookupHeadToLeaf[node][1]
        return headToLeafChain[idxInHeadToLeafChain - k] \
                if idxInHeadToLeafChain - k >= 0 else -1