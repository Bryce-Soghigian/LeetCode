class Solution:
    def remove(self,string):
        
        stack = []
        for char in string:
            if char != "#":
                stack.append(char)
            elif len(stack) != 0:
                stack.pop()
        return "".join(stack)
    def backspaceCompare(self, s: str, t: str) -> bool:
  
        return self.remove(s) == self.remove(t)
                