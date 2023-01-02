class Solution:

    directions = [(-1, 0),(1,0), (0, -1),(0, 1)]
                    
    def in_bounds(self, board, x, y):
        try:
            board[x][y] = board[x][y]
            return True
        except IndexError:
            return False
    
    def dfs(self, board,x,y):
        if not self.in_bounds(board, x, y) or board[x][y] != "O":
            return
        board[x][y] = "S"
        
        for dx, dy in self.directions:
            self.dfs(board, dx + x, dy+ y)
                    
                    
    def solve(self, board):
        if not board or not board[0]:
            return
        

        self.ROWS = len(board)
        self.COLS = len(board[0])

        # Step 1). retrieve all border cells
        from itertools import product
        borders = list(product(range(len(board)), [0, len(board[0])-1])) \
                + list(product([0, len(board)-1], range(len(board[0]))))
        
        #2. Mark All Os connected to border as safe
        for x, y in borders:
            if board[x][y] == "O":
                self.dfs(board, x, y)
        
        
        #3 Revert
        for i in range(len(board)):
            for j in range(len(board[0])):
                if board[i][j] == "O": board[i][j] = "X"
                elif board[i][j] == "S": board[i][j] = "O"