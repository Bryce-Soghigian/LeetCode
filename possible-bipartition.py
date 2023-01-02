class Solution:
    def dfs(self, currentLevel, i, graph, visited):     
        if visited[i] >= 0:
            return (currentLevel - visited[i]) % 2 == 0

        visited[i] = currentLevel
        for child in graph[i]:
            if not self.dfs(currentLevel + 1, child, graph, visited):
                return False
        return True
    def possibleBipartition(self, n, dislikes):

        graph = [[] for i in range(n+1)]
        visited = [-1] * (n+1)
        count = 0
        for dislike in dislikes:
            graph[dislike[0]].append(dislike[1])
            graph[dislike[1]].append(dislike[0])

        for i in range(1, n+1):
            if visited[i] == -1 and  len(graph[i]) > 0:
                if not self.dfs(0, i, graph, visited):
                    return False   
        return True