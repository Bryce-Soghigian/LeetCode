class Solution:
    def shiftGrid(self, grid: List[List[int]], k: int) -> List[List[int]]:
        n, m = len(grid), len(grid[0])
        grid = list(chain(*grid))
        N = len(grid)
        k = k % N
        grid[:] = grid[N-k: N] + grid[:N-k]
        ans = []
        x = 0
        for j in range(n):
            ans.append(grid[x:x+m])
            x += m
        return ans