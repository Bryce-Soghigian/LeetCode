class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        return sorted([nums[i]** 2 for i in range(len(nums))])
        