class Solution:
    def removeOccurrences(self, s: str, part: str) -> str:
        idx = s.find(part)
        if idx == -1:
            return s
        else:
            return self.removeOccurrences(s[:idx]+s[idx+len(part):], part)