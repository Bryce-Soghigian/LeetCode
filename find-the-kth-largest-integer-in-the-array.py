class Solution:
    def kthLargestNumber(self, nums: List[str], k: int) -> str:
        nums = [int(x) for x in nums]
        nums.sort()
        return str(nums[-k])