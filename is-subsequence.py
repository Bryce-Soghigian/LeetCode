class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        if len(s) == 0:
            return True
        
        s_pt = 0
        
        for i,val in enumerate(t):
            if val == s[s_pt]:
                s_pt += 1
            if s_pt > len(s) - 1:
                return True
        return False
        
        