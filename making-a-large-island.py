# class Solution:
#     def _in_bounds(self, grid, x, y):
#         if x < 0 or y < 0 or x > len(grid) - 1 or y > len(grid[0]) -1:
#             return False
#         return True
    
#     def _get_area(self, graph, x, y):
#         """
#         bfs will get teh area of a given island
#         """
#         area = 0
#         queue = deque([(x,y)])
#         seen = set()
#         while queue:
            
#             dx,dy = queue.popleft()
#             seen.add((dx,dy))
#             area += 1
            
#             for i, j in self._directions_iter(dx,dy):
#                 if self._in_bounds(graph,i, j) and graph[i][j] == 1 and (i,j) not in seen:
#                     queue.append((i,j))
#         return area
#     def _set_area(self, graph,hashtable, x, y, value):
#         print(value)
#         queue = deque([(x,y)])
#         seen = set()
#         while queue:
            
#             dx,dy = queue.popleft()
#             graph[dx][dy] = 2
#             hashtable[(dx,dy)] = value
#             seen.add((dx,dy))
            
#             for i, j in self._directions_iter(dx,dy):
#                 if self._in_bounds(graph,i, j) and graph[i][j] == 1 and (i,j) not in seen:
#                     queue.append((i,j))
        
        
#     def _directions_iter(self, x,y):
#         return [(x+1, y), (x-1, y), (x, y+1), (x, y-1)]
    
    
#     def _get_neighbor_sum(self, grid, i, j, area_map):
#         """
#         1 0
#         0 1
#         """
#         seen = set()
        
#         for x,y in self._directions_iter(i, j):
            
#             if self._in_bounds(grid, x,y):
#                 seen.add((x,y))
#         output = 0

#         for area in seen:
#             output += area_map[area][1]
        
#         return output
    
#     def calculate_all_area(self, graph):
#         group_id = 0
#         hashtable = {}
#         for i in range(len(graph)):
#             for j in range(len(graph[0])):
#                 if graph[i][j] == 1:
#                     area = self._get_area(graph, i, j)
#                     self._set_area(graph, hashtable, i, j, (group_id, area))
#                 elif graph[i][j] == 0:
#                     hashtable[(i,j)] = (group_id, 0)
#                 group_id += 1
                    
#         return hashtable
                    
#     def largestIsland(self, grid: List[List[int]]) -> int:
        
#         """
#         1. Find area for all nodes
#         2. Find max area
        
#         (x,y) -> (island_group, area)
        
        
#         """
#         if len(grid) == 0:
#             return 0
        
#         area_map = self.calculate_all_area(grid)
        
#         seen_zero = False
#         max_so_far = 0
#         for i in range(len(grid)):
#             for j in range(len(grid[0])):
#                 if grid[i][j] == 0:
#                     # iterate through the neighbor cells, take 1 from each group
#                     large_island_area = 1 + self._get_neighbor_sum(grid, i, j, area_map)

#                     max_so_far = max(large_island_area, max_so_far)
#                     seen_zero = True
#         if not seen_zero:
#             return area_map[(0,0)][0]
        
#         return max_so_far
class dset:
    
    def __init__(self): #parent, rank and size are hash maps. This data structure requires the nodes to be hashable objects suporting an equality operator, e.g. tuples, strings, integers.
        self.rank=dict()
        self.parent=dict()
        self.size=dict()
    
    def add(self,node): #adds a hashable node object to disjoint set
        if node in self.parent:
            return
        self.parent[node]=node
        self.rank[node]=0
        self.size[node]=1
    
    def find(self,node): #returns the root of node. Then compresses the path so all path nodes point to parent.
        output=node
        parent=self.parent
        
        while parent[output]!=output:
            output=parent[output]
        
        path_node=output
        while parent[path_node]!=output:
            temp=path_node
            path_node=parent[path_node]
            parent[temp]=output
        
        return output
    
    def union(self,node,other):
        """
        unions node and other by rank:
        
        replace node,other by their roots.
        If both equal, return (do nothing)
        re-order node,other so node has greater than or equal rank
        set parent of other to node
        if node and other have equal rank, increase rank of node by 1
        add size of other to size of node
        """
        parent,rank,size=self.parent,self.rank,self.size
        node,other=self.find(node),self.find(other)
        if node==other:
            return
        
        if rank[other]>rank[node]:
            node,other=other,node
        
        parent[other]=node
        if rank[node]==rank[other]:
            rank[node]+=1
        
        size[node]+=size[other]

class Solution:
    def largestIsland(self, grid: List[List[int]]) -> int:
        """
        using disjoint set data structure with size:
        
        initialize distjoint set
        add tuples (r,c) of grid containing 1 to disjoint set
        
        loop over tuples (r,c) of grid:
            if grid at (r,c) is 1:
                for any neighbor tuple (new_r,new_c) lies in grid:
                    if grid at (new_r,new_c) is also 1:
                        union (r,c) with (new_r,new_c)
        
        initialize output at 0
        
        loop over tuples (r,c) of grid:
            if grid at (r,c) is 1:
                update output with size of root of (r,c)
            elif grid at (r,c) is 0:
                find the components associated to each neighbor by querying each neighboring co-ordinate and adding the root to a set.
                update output with 1+sum of size of roots
        
        return output
        """
        R,C=len(grid),len(grid[0])
        islands=dset()
        
        def neighbors(r,c):
            output=set()
            for new_r,new_c in [(r+1,c),(r-1,c),(r,c+1),(r,c-1)]:
                if 0<=new_r<=R-1 and 0<=new_c<=C-1:
                    output.add((new_r,new_c))
            return output
        
        for r in range(R):
            for c in range(C):
                if grid[r][c]==1:
                    islands.add((r,c))
        
        for r in range(R):
            for c in range(C):
                if grid[r][c]==1:
                    for new_r,new_c in neighbors(r,c):
                        if grid[new_r][new_c]==1:
                            islands.union((r,c),(new_r,new_c))
        
        output=0
        for r in range(R):
            for c in range(C):
                if grid[r][c]==1:
                    output=max(output,islands.size[islands.find((r,c))])
                elif grid[r][c]==0:
                    roots=set()
                    for new_r,new_c in neighbors(r,c):
                        if grid[new_r][new_c]==1:
                            roots.add(islands.find((new_r,new_c)))
                    cache=sum([1]+[islands.size[root] for root in roots])
                    output=max(output,cache)
        
        return output