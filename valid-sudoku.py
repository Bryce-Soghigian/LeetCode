class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        """
        Valid Sudoku
        
        
        
        row_sets  = [] # store rows 1-9
        col_sets = [] # store cols 1-9
        box_sets = [] # store box 1-9
        
        Check 
        
        
        
        0 < 2, 3 < 5, 
        """
        N = 9
        box_size = 3
        row_sets = [set() for _ in range(N)]
        col_sets = [set() for _ in range(N)]
        box_sets = [set() for _ in range(N)]
        
        
        for row in range(len(board)):
            
            for col in range(len(board)):
                curr_node = board[row][col]
                box_val = (row//box_size) * box_size + col // box_size

                
                if curr_node != ".":
                    
                    if curr_node in row_sets[row]:
                        return False
                    if curr_node in col_sets[col]:
                        return False
                    if curr_node in box_sets[box_val]:
                        return False
                    
                row_sets[row].add(board[row][col])
                col_sets[col].add(board[row][col])
                box_sets[box_val].add(curr_node)
        return True