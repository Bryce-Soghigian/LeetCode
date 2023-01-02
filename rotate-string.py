class Solution:
    def rotateString(self, s: str, goal: str) -> bool:
        
        """
        We can shift the string s.length number of times and see if any of them are equal
        
        
        
        """
        
        k = len(s)
        shifted = goal
        while k != -1:
            if s == shifted:
                return True
            else:
                
                curr = shifted[0]
                last = shifted[1:]
                shifted = (last + curr)
      
            k -= 1
        return False
                