class UnionFind:
    def __init__(self, n):
        # Each graph node is its
        self.rank = [0 for i in range(n)]
        self.parent = [i for i in range(n)]
    
    def union(self, a, b):
        """
        changes all of the b nodes to join a's group
        """
        root_a = self.find(a)
        root_b = self.find(b)
        
        if root_a == root_b:
            return # dont join these bois the same
        
        if self.rank[root_a] >= self.rank[root_b]:
            self.parent[root_b] = root_a
            self.rank[root_a] = max(self.rank[root_b] + 1, self.rank[root_a])
        else:
            self.parent[root_a] = root_b
            self.rank[root_b] = max(self.rank[root_a] + 1, self.rank[root_b])
            
    def same(self, a, b):
        return self.find(a) == self.find(b)
    
    def find(self, node):
        """
        Return true or false if these two nodes are connected
        """
        parent = self.parent[node]
        while parent != node:
            node = parent
            parent = self.parent[node]
        return parent


class Solution:
    def countComponents(self, n: int, edges: List[List[int]]) -> int:
        
        uf = UnionFind(n)
        
        for edge in edges:
            uf.union(edge[0], edge[1])
        
        sets = set()
        
        for node in range(n):
            sets.add(uf.find(node))
        return len(sets)