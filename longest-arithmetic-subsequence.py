class Solution:
    def longestArithSeqLength(self, nums: List[int]) -> int:
        dp = [Counter()]
        longestArithmeticSequence = 1
        for i in range(1,len(nums)):
            dp.append(Counter())
            for j in range(i):
                difference = nums[i]-nums[j]
                dp[i][difference] = 1 + dp[j][difference]
                longestArithmeticSequence = max(longestArithmeticSequence, dp[i][difference])
        return longestArithmeticSequence + 1