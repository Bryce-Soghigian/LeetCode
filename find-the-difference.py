class Solution:
    def findTheDifference(self, s: str, t: str) -> str:
        t = Counter(t)
        
        for char in s:
            t[char] -= 1
            if t[char] <= 0:
                del t[char]
        
        return list(t.keys())[0]