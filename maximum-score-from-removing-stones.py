class Solution:
    def maximumScore(self, a: int, b: int, c: int) -> int:
        ans = 0
        ar = sorted([a, b, c])
        while ar[0]!=0:
            ar.sort()
            ar[0]-=1
            ar[-1]-=1
            ans+=1
        return ans+min(ar[1], ar[-1])