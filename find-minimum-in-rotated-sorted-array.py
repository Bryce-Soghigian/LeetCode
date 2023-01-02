class Solution:
    def findMin(self, nums: List[int]) -> int:
        """
       [4,5,6,7,0,1,2] 
       s.     m.    e
              s
              7 0 1 2
              s. m  e
              
              7 0
                m
              s
                e
        
        """
        
        start = 0
        end = len(nums) - 1
        
        while nums[start] > nums[end]:
            mid = (start + end) // 2
            if nums[mid] < nums[end]:
                end = mid
            else:
                start = mid + 1
        return nums[start]
                
        
        