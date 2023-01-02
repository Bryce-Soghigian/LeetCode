class Solution:
    def all_upper(self, word, start):
        for i in range(start, len(word)):
            if word[i].islower():
                return False 
        return True

    def all_lower(self, word, start):
        for i in range(start, len(word)):
            if word[i].isupper():
                return False
        return True

    def detectCapitalUse(self, word: str) -> bool:
        if len(word) <= 1: 
            return True
        firstChar = word[0]
        if firstChar.isupper():
            return self.all_lower(word, 1) or self.all_upper(word, 1)
        return self.all_lower(word, 1)