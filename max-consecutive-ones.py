class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        
        
        
        count = 0
        
        max_so_far = 0
        for i in range(len(nums)):
            if nums[i] == 0:
                
                count = 0
            else:
                count += 1
            max_so_far = max(count, max_so_far)     
        return max_so_far