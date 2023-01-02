class Solution:
    def is_pal(self, string):
        
        start = 0
        end = len(string)-1
        while start <= end:
            
            if string[start] != string[end]:
                return False
            start += 1
            end -= 1
        return True
    def firstPalindrome(self, words: List[str]) -> str:
        
        for word in words:
            if self.is_pal(word):
                return word
        return ""
        
