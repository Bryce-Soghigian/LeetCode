class Solution:
    def uncommonFromSentences(self, s1: str, s2: str) -> List[str]:
        s_count = Counter(s1.split(" "))
        s2_count = Counter(s2.split(" "))
        
        
        s_u_count = {key for key, value in s_count.items() if value == 1}
        s2_u_count = {key for key, value in s2_count.items() if value == 1}
        
        for char in s_u_count:
            if char not in s2_count:
                yield char
        
        for char in s2_u_count:
            if char not in s_count:
                yield char