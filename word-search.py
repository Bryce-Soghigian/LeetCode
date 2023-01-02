# class Solution:
#     def in_bounds(self,x,y,grid):
#         if x< 0 or y < 0 or x >=len(grid)-1 or y>= len(grid[x])-1:
#             return False
#         return True
#     def bfs(self,x,y,grid,word):
#         directions = [[0,1],[0,-1],[-1,0],[1,0]]
#         queue = deque()
#         seen = set()
#         queue.append((x,y,0))
#         while queue:
#             x,y,z = queue.popleft()
#             print(x,y,word[z])
#             if grid[x][y] == word[z] and z == len(word) - 1:
#                 return True
                
                
#             for dx, dy in directions:

#                 if self.in_bounds(x + dx,y + dy,grid) and word[z+1] == grid[x+dx][y+dy] and (x+dx, y+dy) not in seen:
#                     seen.add((x+dx, y+dy))
#                     queue.append((x + dx,y + dy,z+1))
                    
                
            
            
#     def exist(self, board: List[List[str]], word: str) -> bool:
#         """
        
#         We have a target char so we can either bfs traversing level by level 
        
        
#         (x,y,level)
        
        
        
#         queue = deque()
        
#         while queue:
#             x,y,level = queue.popleft()
            
#             # then essentally we loop through teh different directions we can go
#             if one of these directions is a match(aka the next word in our search)
#             we add it to the queue else we dont
            
            
        
        
#         only start the bfs when we reach a the first letter
        
#         """
#         for i in range(len(board)):
#             for j in range(len(board)):
#                 if board[i][j] == word[0]:
#                     if self.bfs(i,j,board,word):
#                         return True
#         return False
from collections import Counter
from random import randint

class Solution:
             
    def exist(self, board: List[List[str]], word: str) -> bool:
        
        word_len = len(word)
        if word_len == 0:
            return True
    
        rows, cols = len(board), len(board[0])
        
        # Prune #1
        if word_len > rows * cols:
            return False
        
        # Prune #2
        word_chars = Counter(word)
        matrix_chars = Counter()
        for r in range(rows):
            for c in range(cols):
                matrix_chars[board[r][c]] += 1
        
        if len(word_chars - matrix_chars) > 0:
            return False
        
        def neighbours_of(x, y):
            for dx, dy in {(0, 1), (1, 0), (-1, 0), (0, -1)}:
                cx, cy = dx + x, dy + y
                if cx < 0 or cy < 0 or cx >= rows or cy >= cols:
                    continue
                yield (cx, cy)
        
        def dfs(row, col, visited, idx = 0):
            
            visited.add((row, col))
            
            if idx + 1 == word_len:
                return True
            
            next_char = word[idx + 1]
            for nrow, ncol in neighbours_of(row, col):
                if board[nrow][ncol] == next_char:
                    if (nrow, ncol) in visited:
                        continue
                        
                    if dfs(nrow, ncol, visited, idx + 1):
                        return True
            
            visited.remove((row, col))
            
            return False
        
        first_char = word[0]
        for r in range(rows):
            for c in range(cols):
                if board[r][c] == first_char and dfs(r, c, set()):
                    return True
        
        return False
                
                     