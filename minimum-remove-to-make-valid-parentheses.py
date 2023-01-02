class Solution:
    def minRemoveToMakeValid(self, s: str) -> str:
        """
        
        
        
        stack = []
        1. Find the indexes that do not match
        Everytime we store something we want to store its index as well
        
        ("(",i)
        2. Single Loop adding all of the valid indices
        
        """
        
        stack = []
        braces = "()"
        ind = set()
        for i,char in enumerate(s):
            if stack and stack[-1][0] == "(" and char == ")":
                ind.remove(stack[-1][1])
                stack.pop()
            elif char in braces:
                ind.add(i)
                stack.append((char,i))
        
        
        output = ""
        
        for i in range(len(s)):
            if i not in ind:
                output += s[i]
        return output
        