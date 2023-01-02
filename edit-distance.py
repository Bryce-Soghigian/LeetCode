class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        import math
        A, B = len(word1), len(word2)
        dp = [[math.inf] * (B + 1) for i in range(A + 1)]

        for j in range(B + 1):
            dp[A][j] = B - j

        for i in range(A + 1):
            dp[i][B] = A - i

        for i in range(A - 1, -1, -1):
            for j in range(B - 1, -1, -1):
                if word1[i] == word2[j]:
                    dp[i][j] = dp[i + 1][j + 1]
                if word1[i] != word2[j]:
                    dp[i][j] = 1 + min(dp[i + 1][j + 1], dp[i + 1][j], dp[i][j + 1])

        return dp[0][0]