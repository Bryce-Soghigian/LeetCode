class Solution:
    def isToeplitzMatrix(self, matrix: List[List[int]]) -> bool:
        """
        Iterate through the matrix 
        Starting at (1,1)
        and check that row-1, col-1 is the wsame value if its not return false
        else 
        return true
        
        """
        
        for i in range(1, len(matrix)):
            for j in range(1, len(matrix[i])):
                if matrix[i][j] != matrix[i-1][j-1]:
                    return False
        return True
                