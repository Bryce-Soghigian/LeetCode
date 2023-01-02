class Solution:
    def areOccurrencesEqual(self, s: str) -> bool:
        
        hash = {}
        for char in s:
            if char in hash:
                hash[char] += 1
            else:
                hash[char] = 1
        
        vals = list(hash.values())
        isUni = vals[0]
        
        for val in vals:
            if val != isUni:
                return False
        return True
        