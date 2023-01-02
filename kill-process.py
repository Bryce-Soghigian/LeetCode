class Solution:
    def killProcess(self, pid: List[int], ppid: List[int], kill: int) -> List[int]:
        """
        1. Build the graph
        pid[i]
        2. BFS from the kill point and delete all the nodes from the list
        

        
        graph key : set()
        """
        graph = defaultdict(set)
        
        for i, process in enumerate(ppid):
            graph[process].add(pid[i])
        
        if kill is None:
            return []
        queue = deque([kill])
        
        while queue:
            curr = queue.popleft()
            if curr is not None: 
                yield curr
            
            for child in graph[curr]:
                queue.append(child)
            
            
        