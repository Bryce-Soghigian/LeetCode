class Solution:
    def countLetters(self, s: str) -> int:
        if len(s) == 0:
            return 0
        """
        Two approaches
        
        1. Brute Force
        
        2. Sliding window
        """
        count = 0
        
        for i in range(len(s)):
            seen = set()
            for j in range(i,len(s)):
                seen.add(s[j])
                if len(seen) == 1:
                    count += 1
        return count