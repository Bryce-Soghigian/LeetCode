class Solution:
    def dfs(self,matrix,i,j):
        if not self.in_bounds(matrix,i,j):
            return 0
        directions = [[1,0],[-1,0],[0,-1],[0,1]]
        matrix[i][j] = 0
        area = 1
        for direction in directions:
            di ,dj = direction[0], direction[1]
            area += self.dfs(matrix, i + di, j + dj)
        return area
            
        
    def in_bounds(self,matrix,i,j):
        if i < 0 or j < 0 or j >= len(matrix[0]) or i >= len(matrix) or matrix[i][j] == 0:
            return False
        else:
            return True
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        """
        global_max = 0
        local_max = 0
        
        
        1. We need to write a dfs that counts and visits the island
            def in_bounds(matrix,i,j):
            # if this reutrns false we return 
            else we traverse
            directions = [[1,0]]
            def dfs(matrix, i, j):
            if not in_bounds(matrix,i,j):
                reutrn 
            else:
            matrix[i][j] = 0
            area = 1
            for direction in directions:
                area += dfs(for current dirs)
            return max
            
        
        2. Loop through and when we see a 1 we dfs and compore via kanades logic
        
        # Two loops 
        
        if grid[i][j] == 1:
            global_max = max(dfs(grid,i,j),global_max)
            
        """
        global_max = 0
        for i in range(len(grid)):
            for j in range(len(grid[i])):
                if grid[i][j] == 1:
                    global_max = max(global_max, self.dfs(grid,i,j))
        return global_max
                    