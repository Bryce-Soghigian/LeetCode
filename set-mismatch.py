class Solution:
    def findErrorNums(self, nums: list[int]) -> list[int]:

        total_sum, unique_sum = sum(nums), sum(set(nums))
        n = len(nums)
        return [total_sum - unique_sum, (n * (n + 1)//2)-unique_sum]