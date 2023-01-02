class Solution:
    def maxArea(self, height: List[int]) -> int:
        """
        
        [1,8,6,2,5,4,8,3,7]
        
         s               
                         e
                         
        
        
        
        1. Brute Force
        O(n^2)
        move forward our two pointers try every combo and return the max combo
        
        for i in range(height):
        
        for j in range(height)L
        max_so = max(ma_so, min(height[i], height[j]) * j-i)
        
        
        O(n) TC
        O(3) SC
        2. 
        
        1. Start pointers on both ends of our list
        2. if s <= e:
           move start forward
           calculate new max_area
          else:
          move in end calculate new max area
           height = min(height[start], height[end])
           max_so_far = max(height*width, max_so_far)
           
        
        
        """
        start = 0
        end = len(height) - 1
        max_so_far = 0
        while start <= end:


            
            min_height = min(height[start], height[end])
            width = end - start
            max_so_far = max(max_so_far, min_height * width)
            if height[start] < height[end]:
                start += 1
            else:
                end -= 1
        return max_so_far
            
        