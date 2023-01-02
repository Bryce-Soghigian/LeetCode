class Solution:
    def _swap(self, s, start, end):
    
        
        while start <= end:
            
            s[start], s[end] = s[end], s[start]
            start += 1
            end -= 1
            
            
    def reverseWords(self, s: List[str]) -> None:

        
        self._swap(s, 0, len(s) -1)
        
        start = 0
        s.append(" ")
        for i, char in enumerate(s):
            
            if char == " ":
                self._swap(s,start, i-1)
                start = i+1
        s.pop()
        
                