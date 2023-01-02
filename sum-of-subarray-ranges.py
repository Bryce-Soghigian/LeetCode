class Solution:
    def subArrayRanges(self, nums: List[int]) -> int:
        sum_of_ranges = 0
        for i in range(len(nums)):
            minimum = float('inf')
            maximum = float('-inf')
            for j in range(i, len(nums)):
                minimum = min(nums[j], minimum)
                maximum = max(nums[j], maximum)
                sum_of_ranges += (maximum - minimum)
                
        return sum_of_ranges