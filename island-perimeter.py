class Solution:
    def islandPerimeter(self, grid: List[List[int]]) -> int:
        perimeter = 0
        for r in range(len(grid)):
            for c in range(len(grid[0])):
                if grid[r][c] == 1:
                    perimeter += self.numSides(grid, r, c)
        
        return perimeter    
    def numSides(self, grid, row, col):
        sides = [[-1,0], [0,1], [0, -1], [1, 0]]
        ans = 0
        for side in sides:
            r = row + side[0]
            c = col + side[1]
            ans += 1
            if r >= 0 and r < len(grid) and c >= 0 and c < len(grid[0]):
                if grid[r][c] == 1:
                    ans -= 1
        return ans