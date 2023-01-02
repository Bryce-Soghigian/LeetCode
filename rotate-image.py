class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        
        
        iterate through list col by col
        store each iterated piece in an array
        assign matrix[i] = rotatedPiece[i]
        
        Temp = 1
        [[1,2,3]
        ,[4,5,6]
        ,[7,8,9]]
        
        [[7,2,1]
        ,[4,5,6]
        ,[9,8,3]]
        
        
        [[7,4,1]
        ,[8,5,2]
        ,[9,6,3]]
        
        """
        length = len(matrix) -1
        for i in range(length):
            for j in range(i, length - i):
                temp = matrix[i][j]
                matrix[i][j] = matrix[length - j][i]
                matrix[length-j][i] = matrix[length-i][length-j]
                matrix[length - i ][length - j] = matrix[j][length-i]
                matrix[j][length - i] = temp
        