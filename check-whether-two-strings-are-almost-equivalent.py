class Solution:
    def checkAlmostEquivalent(self, word1: str, word2: str) -> bool:
        """
        For word1 we want to check to see if word2 has the same chars
        
        """
        
        

        words_1 = Counter(word1)
        words_2 = Counter(word2)
        
        for key,value in words_1.items():
            if words_2[key] != value and abs(words_2[key] - value) > 3:
                return False
        
        for key, value in words_2.items():
            if words_1[key] != value and abs(words_1[key] - value) > 3:
                return False
        return True
               
                
        