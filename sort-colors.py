class Solution:
    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        
        
        Approaches
        
        1. Bucket sort
        
            [0,0,0]
             0,1,2
             
        Iterate through array
        Count each color
        Iterate through array and replace colors
        
        
        2. Three Pointers
        
        
        [2,0,2, 1,1, 0]
                     blue
         white       
                      
        """
        
        
        buckets = [0,0,0]
        
        for val in nums:
            buckets[val] += 1
        
        
        pt = 0
        
        for i in range(len(nums)):
            
            while buckets[pt] == 0:
                pt += 1
            nums[i] = pt
            buckets[pt] -= 1
            
            
            
        
        
        