class Solution:
    indegree = []
    def _build_graph_info(self, edges, n):
        """
        Builds an adjency list from a set of edges and a number of nodes
        
        """
        self.indegree = [0 for _ in range(n)]
        graph = [[] for _ in range(n)]
        for source, dest in edges:
            graph[source].append(dest)
            self.indegree[dest] += 1
            
        return graph

    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        """
        0. Build the graph
        1. count the indegree of each node
        2. add the non-zero indegrees to a queue
        3. while that queue has elements we want to prune the tree pushing the indegree nodes of zero into output, 
        3a. for each indegree node decrement its children
        if
        """
        # build graph and indegree mapping
        graph = self._build_graph_info(prerequisites, numCourses)
        
        zero_indegree = deque()
        for i, degree in enumerate(self.indegree):
            if degree == 0:
                zero_indegree.append(i)
        
        output = []
        while zero_indegree:
            size = len(zero_indegree)
            for _ in range(size):
                curr = zero_indegree.popleft()
                output.append(curr)
                for child in graph[curr]:
                    # decrement indegree, if indegree == 0 after decrment add to queue
                    self.indegree[child] -= 1
                    if self.indegree[child] == 0:
                        zero_indegree.append(child)
        return len(output) == numCourses
        
        