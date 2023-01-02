class Solution:
    def is_palindrome(self, string):
        
        start = 0
        end = len(string) - 1
        while start <= end:
            if string[start] != string[end]:
                return False
            start += 1
            end -= 1
        return True
        
    def countSubstrings(self, s: str) -> int:
        """
        
        
        abcd
        
        
        
        """
        count = 0
        for i in range(len(s)):
            substring = ""
            for j in range(i, len(s)):
                substring += s[j]
                if self.is_palindrome(substring):
                    count += 1
        
        return count