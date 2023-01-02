class Solution:
    def reverseOnlyLetters(self, s: str) -> str:
        
        
        
        stack = []
        
        for char in s:
            if char.isalpha():
                stack.append(char)
                
        s = [k for k in s]
        for i in range(len(s)):
            if s[i].isalpha():
                new = stack.pop()
                s[i] = new
        return "".join(s)