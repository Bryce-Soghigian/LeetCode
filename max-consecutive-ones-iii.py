class Solution:
    def longestOnes(self, nums: List[int], k: int) -> int:
        start = 0
        last = 0
        maximum = 0
        while last < len(nums):
            if nums[last] == 0:
                k-=1
                
            while k < 0:
                if nums[start] == 0:
                    k += 1
                start += 1
            
            maximum = max(maximum, last - start + 1)
            last += 1
                
        return maximum
            
        