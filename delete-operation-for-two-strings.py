class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        "lev idea"
        dp = [i for i in range(len(word2) + 1)]
        for i in range(1, len(word1) + 1):
            # prev is used to store dp[i - 1][j - 1]
            prev, dp[0] = i - 1, i
            for j in range(1, len(dp)):
                if word1[i - 1] == word2[j - 1]:
                    prev, dp[j] = dp[j], prev
                else:
                    prev, dp[j] = dp[j], 1 + min(dp[j], dp[j - 1])
        return dp[-1]