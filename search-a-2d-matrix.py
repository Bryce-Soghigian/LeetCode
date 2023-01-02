class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        def dfs(x, y):
            if not(0 <= x < n) or not(0 <= y < m):
                return False
            if (x,y) in lookup:
                return False
            lookup.add((x,y))
            if matrix[x][y] == target:
                return True
            if matrix[x][y] < target:
                down = dfs(x+1, y)
                right = dfs(x, y+1)
                return down or right
            else:
                top = dfs(x-1, y)
                left = dfs(x, y-1)
                return top or left
        
        n, m = len(matrix), len(matrix[0])
        lookup = set()
        return dfs(0,0)
