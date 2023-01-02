class Solution:
    def checkIfPangram(self, sentence: str) -> bool:
        letters = set()
        for i in range(97, 97 + 26):
            letters.add(chr(i))

        for char in sentence:
            if char in letters:
                letters.discard(char)
        
        return len(letters) == 0

        
