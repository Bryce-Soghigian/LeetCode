class Solution:
    def numOfStrings(self, patterns: List[str], word: str) -> int:
        
        
        seen = set()
        
        
        for i in range(len(word)):
            sub = ""
            for j in range(i,len(word)):
                sub += word[j]
                seen.add(sub)
        
        count = 0
        for word in patterns:
            if word in seen:
                count += 1
                
        return count
        