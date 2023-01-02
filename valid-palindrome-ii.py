class Solution:
    
    def validPalindrome(self, s: str) -> bool:
        
        def checkPalindrome(s: str, err: int = 1):
            left = 0
            right = len(s)-1
            while left < right:
                if s[left] != s[right] and err == 1:
                    return checkPalindrome(s[left+1:right+1], err-1) or checkPalindrome(s[left:right], err-1)
                elif s[left] != s[right] and err == 0:
                    return False
                left += 1
                right -= 1
            return True
        
        return checkPalindrome(s, 1)