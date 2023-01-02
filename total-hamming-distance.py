class Solution:
    def totalHammingDistance(self, nums: List[int]) -> int:
        count, result, N = defaultdict(int), 0, len(nums)
        for num in nums:
            for pos in range(31):
                if num & 1: count[pos] += 1
                num >>= 1
                if not num: break

        for pos in range(31):
            result += count[pos]*(N-count[pos])

        return result