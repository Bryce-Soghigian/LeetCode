class Solution:
    def minSteps(self, s: str, t: str) -> int:
        """
        Count the frequency of characters and compare the differences
        
        
        
        B: 2
        a: 1
        
        
        a: 2
        b: 1
        
        The difference is 1
        
        
        """
        s, t = Counter(s), Counter(t)
        res = 0
        for i in t:
            if t[i] >= s[i]:
                res += t[i] - s[i]
        
        return res
        
        
        