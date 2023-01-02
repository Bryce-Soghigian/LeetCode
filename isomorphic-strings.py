class Solution:
    def is_iso(self, s,t):
        hashmap = {}
        
        
        for i, val in enumerate(s):
            if not val in hashmap:
                hashmap[val] = t[i]
            else:
                if hashmap[val] != t[i]:
                    return False
        return True
    def isIsomorphic(self, s: str, t: str) -> bool:
        return self.is_iso(s,t) and self.is_iso(t,s)
        