class Solution:
    def shortestPathBinaryMatrix(self, grid: List[List[int]]) -> int:
        
        M = len(grid)
        N = len(grid[0])
        
        if grid[0][0] != 0 or grid[-1][-1] != 0:
            return -1

        queue=[(0,0)]
        grid[0][0] = 1
        nextQueue = []
        count = 1
        while queue:
            while queue:
                left,right = queue.pop(0)
                
                if left == M-1 and right == N-1:
                    return count
                lst = [(left-1,right-1),(left-1,right),(left-1,right+1),(left,right+1),(left+1,right+1),(left+1,right),(left+1,right-1),(left,right-1)]
                
                for r,c in lst:
                    if r==M-1 and c ==N-1:
                        return count+1
                    if r > -1 and r < M and c > -1 and c < N:
                        if grid[r][c] == 0:
                            nextQueue.append((r,c))
                            grid[r][c] = 1
                            
            count += 1
            queue = nextQueue
            nextQueue = []
            
        return -1
        
        
        