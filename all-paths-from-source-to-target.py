class Solution:
    def allPathsSourceTarget(self, graph: List[List[int]]) -> List[List[int]]:
        if len(graph) == 0: return graph
        paths = []
        

        queue = deque([(0,[])])

        while queue:
            curr, seen = queue.popleft()
            seen.append(curr)
            if curr == len(graph) - 1:
                paths.append(seen)

            for child in graph[curr]:
                queue.append((child, seen[:]))
        return paths

