class Solution:
    def search(self, nums: List[int], target: int) -> int:
        """
        1. Find the min index
        2. Binary search on both halfs
        """
        for i,val in enumerate(nums):
            if val == target:
                return i
        return -1