class Solution:
    def dfs(self,matrix,x,y):
        # loop through row you want to set to zero
        for i in range(len(matrix[x])):
            matrix[x][i] = 0
        # loop through column and set to zero
        for i in range(len(matrix)):
            matrix[i][y] = 0
        print(matrix)
    def setZeroes(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        [1,0,0]
        [1,0,0]
        [1,0,0]
        """
        matrix_copy = copy.deepcopy(matrix)
        for i in range(len(matrix)):
            for j in range(len(matrix[i])):
                if matrix_copy[i][j] == 0:
                    self.dfs(matrix,i,j)
        