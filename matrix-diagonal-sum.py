# class Solution:
#     def _get_diagonal_sum(self, mat, start, direction):
        
#         summation = 0
#         x,y = start
#         for _ in range(len(mat)):
#             summation += mat[x][y]
#             print(summation)
#             x += direction[0]
#             y += direction[1]
#         return summation
#     def diagonalSum(self, mat: List[List[int]]) -> int:
#         """
#         We want to traverse the diagonals
        
#         [[7,3,1,9]
#         ,[3,4,6,9]
#         ,[6,9,6,6]
#         ,[9,5,8,5]]
#          primary = 2 + 27 
        
        
#         """
#         mid_index = len(mat) //2
#         primary = self._get_diagonal_sum(mat, (0,0), [1,1])
#         secondary = self._get_diagonal_sum(mat, (len(mat)-1, 0), [-1, -1])
#         if len(mat) % 2 == 0:
#             print(primary, secondary)
#             return primary + secondary
#         return (primary + secondary) - mat[mid_index][mid_index]

class Solution:
    def diagonalSum(self, mat: List[List[int]]) -> int:
        
        n = len(mat)
        
        mid = n // 2
        
        summation = 0
        
        for i in range(n):
            
            # primary diagonal
            summation += mat[i][i]
            
            # secondary diagonal
            summation += mat[n-1-i][i]
            
            
        if n % 2 == 1:
            # remove center element (repeated) on odd side-length case
            summation -= mat[mid][mid]
            
            
        return summation