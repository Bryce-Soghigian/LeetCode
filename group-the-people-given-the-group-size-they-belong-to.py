class Solution:
    def groupThePeople(self, groupSizes: List[int]) -> List[List[int]]:
        
        graph = defaultdict(list)
        for i, group in enumerate(groupSizes):
            graph[group].append(i)
        

        graph = sorted(graph.items(), key=lambda x: x[1])
        
        output = []
        
        for node in graph:
            bucket = []
            
            for i in range(len(node[1])):
                bucket.append(node[1][i]) 
                if node[0] == len(bucket):
                    output.append(bucket)
                    bucket = []
        return output
                 
