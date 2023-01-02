class Solution:
    def smallestRepunitDivByK(self, K: int) -> int:
        remainder = 0
        
        for result in range(1, K+1):
            if (remainder * 10 + 1) % K == 0:
                return result
            remainder = (remainder * 10 + 1) % K

            
            
        return -1