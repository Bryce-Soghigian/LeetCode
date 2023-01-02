from collections import Counter
class Solution:
    def findAnagrams(self, s: str, p: str) -> List[int]:
        if len(s) < len(p):
            return []
        s_count = Counter()
        p_count = Counter(p)
        
        for end in range(len(s)):

            s_count[s[end]] += 1
            if end > len(p)-1:
                if s_count[s[end - len(p)]] == 1:
                    del s_count[s[end - len(p)]]
                else:
                    s_count[s[end - len(p)]] -= 1

            if p_count == s_count:
                yield end - len(p) + 1
        
  