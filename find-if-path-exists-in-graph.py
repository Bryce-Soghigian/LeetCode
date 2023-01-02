class Solution:
    def validPath(self, n, edges, source, destination):
        # Create a queue to store the vertices to be visited
        queue = []
        # Add the source vertex to the queue
        queue.append(source)
        # Initialize the visited array
        visited = [False] * n
        # Mark the source vertex as visited
        visited[source] = True

        # While the queue is not empty
        while queue:
            # Remove the first vertex from the queue
            current_vertex = queue.pop(0)
            # If the current vertex is the destination, return True
            if current_vertex == destination:
                return True
            # Iterate through the edges of the current vertex
            for u, v in edges:
                # If the edge leads to an unvisited vertex, mark it as visited and add it to the queue
                if u == current_vertex and not visited[v]:
                    visited[v] = True
                    queue.append(v)
                elif v == current_vertex and not visited[u]:
                    visited[u] = True
                    queue.append(u)
        # If the queue becomes empty, return False
        return False
