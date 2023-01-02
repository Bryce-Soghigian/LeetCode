class Solution:
    def maxPower(self, s: str) -> int:
        """
        
        The power of a string is the max length of a non empty substring that contains only one 
        unique character
        
        
                leetcode
                    2
                    ee
                    
                    
            
            sliding window
            
            
            leetcode
             s
            
        
            
            
            
        
        """
        
        
        if len(s) == 0:
            return 0
        
        max_substring = 0
        start = 0
        curr = s[start]
        
        while start <= len(s) -1:
            # move forward pointer while its equal to curr
            diff = start
            while start <= len(s) -1 and s[start] == curr:
                start += 1
            if start <= len(s) - 1:
                curr = s[start]
            max_substring = max(max_substring, start - diff)
        return max_substring