class Solution:
    def shuffle(self, nums: List[int], n: int) -> List[int]:
        
        midpoint = n
        i = 0
        while midpoint <= len(nums) - 1:
            yield nums[i]
            yield nums[midpoint]
            i += 1
            midpoint += 1
        