class Solution:
    def chalkReplacer(self, A, k):
        k %= sum(A)
        for i, a in enumerate(A):
            if k < a:
                return i
            k -= a
        return 0
                
        