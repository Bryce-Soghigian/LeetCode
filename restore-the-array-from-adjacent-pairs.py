class Solution:
    def build_list(self,edges):
        graph = {}
        for e,v in edges:
            graph[e] = []
            graph[v] = []
        for e,v in edges:
            graph[e].append(v)
            graph[v].append(e)
        return graph
    def restoreArray(self, adjacentPairs: List[List[int]]) -> List[int]:
        """
        1. Build a graph
        
        
        2. Dfs on that graph to build your array
        
        """
        graph = self.build_list(adjacentPairs)
        start = None
        for k,adj_list in graph.items():
            if len(adj_list) == 1:
                start = k
                break
        seen = set()
        output = []
        def dfs(index):
            if index in seen:
                return
            else:
                seen.add(index)
                output.append(index)
                for child in graph[index]:
                    dfs(child)
        dfs(start)
        return output
        
        
                