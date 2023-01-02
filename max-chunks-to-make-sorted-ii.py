import math
class Solution:
        def maxChunksToSorted(self, arr: List[int]) -> int:
            right_min = [math.inf] * len(arr)
            left_max = arr.copy()
            for i in range(len(arr) - 2, -1, -1):
                right_min[i] = min(arr[i + 1], right_min[i + 1])
            for i in range(1, len(arr)):
                left_max[i] = max(left_max[i - 1], arr[i])
            chunks = 0
            for i in range(len(arr)):
                if right_min[i] >= left_max[i]:
                    chunks += 1
            return chunks
#     def maxChunksToSorted(self, arr: List[int]) -> int:
#         """
#         We use a recursive approach
        
        
#         we want to look at our current array
        
#         loop through our current array until our nums aren't decreasing.
#         if a num is out of place we want to recurse from [curr:] and 
#         check thast array
#         \
        
        
#         exakmple 
        
#         [2,1,5,4,5]
#         num of chunks = 0
        
#         def count_chunks(num_of_chunks,curr_arr)
#         if len(arr) == 0:
#         return count
        
#         """
#         count = 0
#         strictly_decreasing = True
#         def count_chunks(curr_arr):
#             nonlocal count
#             nonlocal strictly_decreasing
#             if len(curr_arr) == 0:
#                 return
#             if len(curr_arr) == 1:
#                 count += 1
#                 return
#             if len(curr_arr) == 2 and strictly_decreasing == False:
#                 if curr_arr[0] > curr_arr[1]:
#                     count += 1
#                 else:
#                     count += 2
#                 return
            
#             for i in range(1,len(curr_arr)):
#                 if curr_arr[i] >= curr_arr[i-1]:
#                     count += 1
#                     strictly_decreasing = False
#                     count_chunks( curr_arr[i:])
#                     break
            
#         count_chunks(arr)
#         if strictly_decreasing == True:
#             return 1
#         return count
        