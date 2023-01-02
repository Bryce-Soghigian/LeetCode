class Solution:

    def shortestDistance(self, grid: List[List[int]]) -> int:
        rows, cols = len(grid), len(grid[0])

        def neighbors(row, col, curr):
            for r_offset, c_offset in ((1, 0), (-1, 0), (0, 1), (0, -1)):
                r, c = row + r_offset, col + c_offset
                if 0 <= r < rows and 0 <= c < cols and grid[r][c] == 0 and curr[r][c] == -1:
                    yield r, c

        buildings = [(r, c) for r in range(rows) for c in range(cols) if grid[r][c] == 1]
        aggregate = [[0] * cols for _ in range(rows)]

        for row, col in buildings:
            curr = [[-1] * cols for _ in range(rows)]
            q = collections.deque([(row, col, 0)])
            while q:
                row, col, dist = q.popleft()
                for r, c in neighbors(row, col, curr):
                    curr[r][c] = dist + 1
                    q.append((r, c, dist + 1))

            aggregate = [[aggregate[r][c] + curr[r][c] if curr[r][c] != -1 and aggregate[r][c] != -1 else -1 for c in range(cols)] for r in range(rows)]

        return min((aggregate[r][c] for r in range(rows) for c in range(cols) if aggregate[r][c] > 0), default=-1)
