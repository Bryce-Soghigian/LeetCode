class Solution:
#     def get_outdegree(self, matrix):
        
#         output = [0] * len(matrix)
        
#         for i in range(len(matrix)):
#             indegree_sum = 0
#             for j in range(len(matrix[i])):
#                 if matrix[i][j] == 1:
#                     indegree_sum += 1
#             output[i] = indegree_sum - 1
        
        
#         return output
#     def get_indegree(self,matrix):
#         outdegree = [0] * len(matrix)
#         for row in range(len(matrix[0])):
#             outdegree_count = 0
#             for col in range(len(matrix)):
#                 if matrix[col][row] == 1:
#                     outdegree_count += 1
            
#             outdegree[row] = outdegree_count - 1
#         return outdegree
                
                
#     def construct_graph(self,n):
        
#         graph = [[0] * n for i in range(n)]
        
#         for i in range(n):
#             for j in range(n):
#                 if knows(i,j):
#                     graph[i][j] = 1
            
#         return graph
        
        
#     def findCelebrity(self, n: int) -> int:
        
#         sheeesh = self.construct_graph(n)
#         indegrees = self.get_indegree(sheeesh)
#         outdegrees = self.get_outdegree(sheeesh)
        
#         for i in range(n):
#             if indegrees[i] == n-1 and outdegrees[i] == 0:
#                 return i 
#         return -1
    def findCelebrity(self, n: int) -> int:
        notcelebs = set()
        # Looks like O(n^2), actually O(N) since we remove somebody every loop
        for c in range(n):
            for d in range(n):
                if c in notcelebs: # C might have been excluded previously
                    break
                if c == d: # Don't check if you know yourself
                    continue
                if knows(d,c): # Check if everyone knows C
                    notcelebs.add(d) # D->C => !celeb(D)
                else:
                    notcelebs.add(c) # D-\->C => !celeb(C)
                    break # No longer care about C, try someone else
                if knows(c,d): # Check that C doesn't know anyone
                    notcelebs.add(c) # C->D => !celeb(C)
                else:
                    notcelebs.add(d)
            if c not in notcelebs: # Completed loop, ∀(D): C-\->D && D->C
                return c
        return -1
