class Solution:
    def in_bounds(self, grid, x, y):
        if x < 0 or y < 0 or len(grid) <= x or len(grid[0]) <= y:
            return False
        return True
    
    def floodFill(self, image: List[List[int]], sr: int, sc: int, newColor: int) -> List[List[int]]:
        origin = image[sr][sc]
        
        queue = deque([(sr,sc)])
        seen = set()
        while queue:
            
            x,y = queue.popleft()
            seen.add((x,y))
            image[x][y] = newColor
            
            for dx,dy in [[x-1, y], [x+1, y], [x, y+1], [x,y-1]]:
                if self.in_bounds(image,dx,dy) and image[dx][dy] == origin and (dx,dy) not in seen:
                    queue.append((dx,dy))
                    
        return image
            
            
        