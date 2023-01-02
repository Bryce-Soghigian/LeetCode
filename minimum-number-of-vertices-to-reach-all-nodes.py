class Solution:
    def findSmallestSetOfVertices(self, n: int, edges: List[List[int]]) -> List[int]:
        graph = [[] for i in range(n)]
        for e in edges:
            graph[e[1]].append(e[0])
        output = []
        for k in range(len(graph)):
            if len(graph[k]) == 0:
                output.append(k)
        return output
        