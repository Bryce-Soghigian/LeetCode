class Solution:
    def findOrder(self, num_courses: int, prerequisites: List[List[int]]) -> List[int]:
        """
        Important Graph Theory Concepts
        
        Indegree
        
        
        Indegree 1 to num_of_courses - 1[0,0,0,0]
        
        
        lets draw the graph 
        
        
        num_c = 4
        pre = [[1,0],[2,0],[3,1],[3,2]]
        
        to take course 1 you must take course 0
        indegrees = [0,1,1,2]
        0, 1, 2 3
        0, 2, 1, 3
        the indegree for 1 is 1
        0:[1, 2]
        1:[]
        2:[]
        3:[1,2]
        
        we can go through the indegrees and find the zero ones
        
        Two approaches
        
        1. Union Find
        
        2. Topological Sort
            Count the Indegree of the nodes
            
            Gradually decrement indegree and keep adding items to the queue until there are no more nodes of indegree of zero
        """
        
        graph = [[] for _ in range(num_courses)]
        indegree = [0 for _ in range(num_courses)]
        
        
        for src, dest in prerequisites:
            graph[dest].append(src)
            indegree[src] += 1
        
        # We essentially want to gradually decrement our indegree for the children of the nodes with an indegree of zero
        
        zeros = [node for node in range(num_courses) if indegree[node] == 0]
        
        output = []
        while zeros:
            
            for _ in zeros:
                # for these nodes we want to decrement their indegree
                node = zeros.pop()

                output.append(node)
                for child in graph[node]:
                    
                    indegree[child] -= 1
                    if indegree[child] == 0:
                        zeros.append(child)
                        

        return output if len(output) == len(indegree) else []
            
                
        