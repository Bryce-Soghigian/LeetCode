class Solution:
    def maxDepth(self, s: str) -> int:
        """
        We want to recurse every time we see a opening parenthesis 
        
        """
        
        
        stack = []
        max_depth_so_far = 0
        
        for char in s:
            if stack and stack[-1] == "(" and char == ")":
                stack.pop()
            elif char in "()":
                stack.append(char)
            
            max_depth_so_far = max(max_depth_so_far, len(stack))
        
        return max_depth_so_far
        