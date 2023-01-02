class Solution:
    @cache
    def numWays(self, steps: int, arrLen: int, pos = 0) -> int:
        if pos < 0 or pos >= arrLen:
            return 0
        if steps == 0:
            return 1 if pos == 0 else 0
        return (self.numWays(steps-1, arrLen, pos-1)+self.numWays(steps-1, arrLen, pos)+self.numWays(steps-1, arrLen, pos+1))%1000000007