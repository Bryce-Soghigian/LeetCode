class Solution:
    directions = [[-1,0],[0,-1],[1,0],[0,1]]
    
    def _in_bounds(self, matrix, x, y):
        if x < 0 or y < 0 or x >= len(matrix) or y >= len(matrix[0]) or matrix[x][y] == "0":
            return False
        return True
        
    def _dfs(self, x,y,matrix):
        if not self._in_bounds(matrix,x,y):
            return
            # mark node as visited
        matrix[x][y] = "0"
        for dx,dy in self.directions:
            self._dfs(dx+ x, dy + y, matrix)
        
    def numIslands(self, grid: List[List[str]]) -> int:
        count = 0
        for i in range(len(grid)):
            for j in range(len(grid[0])):
                if grid[i][j] == "1":
                    self._dfs(i,j,grid)
                    count += 1
        
        return count
        