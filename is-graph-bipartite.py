class Solution:
    def bfs(self,colors,starting_index,starting_color,graph):
        queue = deque()
        queue.append((starting_index,starting_color))
        while queue:
            current = queue.popleft()
            if colors[current[0]] == -1:
                colors[current[0]] = current[1]
            if colors[current[0]] != current[1] and colors[current[0]] != -1:
                return False

                # Loop through the neighbors
            new_color = current[1] ^ 1
            for child in graph[current[0]]:
                if colors[child] == -1:
                    queue.append((child,new_color))
        return True
    def isBipartite(self, graph: List[List[int]]) -> bool:
        """
        Write a BFS
        
        Each bfs we want to give a color to the children and check if they are colored.
        if the color doesn't match teh current color pattern we return false
        
        if we get through all of the nodes we want to reutrn true
        """
        colors = [-1 for i in range(len(graph))]


        curr_color = 1
        for i in range(len(colors)):
            curr_color ^= 1
            if colors[i] == -1:
                bfs_result = self.bfs(colors,i,curr_color,graph)
                if bfs_result == False:
                    return False
        return True
                
                
            