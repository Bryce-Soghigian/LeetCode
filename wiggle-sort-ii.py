class Solution:
    def wiggleSort(self, nums):
        count = [0]*5001
        
        for n in nums:
            count[n]+=1
        
        odd = 1
        even = 0
        for n in range(5000,-1,-1):
            if odd >= len(nums) and even >= len(nums):
                break
                
            if count[n] == 0:
                continue
            
            while count[n] and (odd < len(nums) or even < len(nums)):
                count[n]-=1
                if odd < len(nums):
                    nums[odd] = n
                    odd+=2
                else:
                    nums[even] = n
                    even+=2
            